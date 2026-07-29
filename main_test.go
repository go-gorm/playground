package main

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"io"
	"sync"
	"testing"
	"time"

	"gorm.io/gorm"

	"gorm.io/playground/models"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := models.User{Name: "jinzhu"}

	DB.Create(&user)

	var result models.User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}

// TestExpiredPreparedStatementLeak reproduces a prepared-statement leak in
// GORM's LRU statement store (internal/lru + internal/stmt_store):
//
//  1. lru.Get returns a miss for an expired entry but leaves the entry (and
//     its *sql.Stmt) in the cache.
//  2. Re-executing the same query before the background cleanup goroutine
//     removes that entry takes lru.Add's "existing item" branch, which
//     overwrites ent.Value WITHOUT invoking the onEvict callback.
//  3. The onEvict callback is the only place where cached statements are
//     closed, so the overwritten *sql.Stmt is never closed. It stays open as
//     a database/sql dependency of the *sql.DB until the whole pool is
//     closed, outside the configured max size — one leaked statement per
//     TTL cycle per hot query.
//
// The test prepares a query, lets its entry expire, then re-executes it
// inside the window before cleanup removes the expired entry (a "filler"
// query expiring 50ms later in the same expiration bucket prevents the
// cleanup goroutine from processing that bucket any earlier). It then waits
// until that bucket's cleanup has run and asserts that GORM has closed the
// replaced statement; on current master it has not, and the statement only
// disappears when the entire *sql.DB is closed.
func TestExpiredPreparedStatementLeak(t *testing.T) {
	const ttl = 10 * time.Second
	tick := ttl / 100 // cleanup bucket granularity; one bucket per tick

	demonstrated, timingMisses := false, 0
	for attempt := 1; attempt <= 3 && !demonstrated; attempt++ {
		target := fmt.Sprintf("SELECT %d", attempt*2)
		filler := fmt.Sprintf("SELECT %d", attempt*2+1)

		tracker := &stmtTracker{preparedAt: map[string][]time.Time{}, closedAt: map[string][]time.Time{}}
		sqlDB, err := sql.Open(registerTrackingDriver(tracker), "")
		if err != nil {
			t.Fatal(err)
		}
		sqlDB.SetMaxOpenConns(1)

		start := time.Now()
		pdb := gorm.NewPreparedStmtDB(sqlDB, 100, ttl)

		exec := func(query string) {
			t.Helper()
			if _, err := pdb.ExecContext(context.Background(), query); err != nil {
				t.Fatalf("exec %q: %v", query, err)
			}
		}

		// Align to the next cleanup-tick boundary so that both statements
		// below land in the same expiration bucket.
		now := time.Now()
		sleepUntil(start.Add(now.Sub(start)/tick*tick + tick + time.Millisecond))

		t0 := time.Now()
		exec(target)                 // expires at ~t0+ttl in bucket B
		sleepUntil(t0.Add(tick / 2)) // +50ms, still the same bucket
		exec(filler)                 // expires at ~t0+ttl+50ms; cleanup of B sleeps until then

		// Re-execute the query just after its entry expires, while the
		// expired entry is still waiting for cleanup in bucket B. The entry's
		// ExpiresAt is stamped at cache-insertion time (slightly after t0), so
		// aim 20ms past t0+ttl; cleanup of bucket B cannot run before the
		// filler expires at ~t0+ttl+50ms.
		sleepUntil(t0.Add(ttl + 20*time.Millisecond))
		exec(target)

		// Wait until bucket B's cleanup pass has finished, then check whether
		// GORM closed the overwritten statement. It must not be confused with
		// database/sql closing it later when the whole pool is closed.
		sleepUntil(t0.Add(ttl + 3*tick)) // +300ms, past the filler's expiry
		prepared, closed := tracker.counts(target)
		fillerPrepared, fillerClosed := tracker.counts(filler)
		t.Logf("attempt %d: target prepared=%d closed=%d, filler prepared=%d closed=%d",
			attempt, prepared, closed, fillerPrepared, fillerClosed)

		pdb.Close()
		sqlDB.Close()

		switch {
		case prepared == 2 && closed == 0:
			demonstrated = true
		case prepared < 2:
			timingMisses++
			t.Logf("attempt %d: re-execution hit the cache before expiry, retrying", attempt)
		default:
			t.Logf("attempt %d: overwritten statement was closed by GORM", attempt)
		}
	}

	if demonstrated {
		t.Fatalf("leaked prepared statement: an expired cache entry was overwritten without Close; " +
			"the old *sql.Stmt remained open after all cache cleanup and is only released by closing the whole *sql.DB")
	}
	if timingMisses == 3 {
		t.Skip("could not hit the expiry-to-cleanup window in 3 attempts")
	}
}

// stmtTracker records, per query, when statements were prepared and closed.
type stmtTracker struct {
	mu         sync.Mutex
	preparedAt map[string][]time.Time
	closedAt   map[string][]time.Time
}

func (t *stmtTracker) addPrepared(query string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.preparedAt[query] = append(t.preparedAt[query], time.Now())
}

func (t *stmtTracker) addClosed(query string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.closedAt[query] = append(t.closedAt[query], time.Now())
}

func (t *stmtTracker) counts(query string) (prepared, closed int) {
	t.mu.Lock()
	defer t.mu.Unlock()
	return len(t.preparedAt[query]), len(t.closedAt[query])
}

var trackingDrivers sync.Map // driver name -> *stmtTracker

func registerTrackingDriver(tracker *stmtTracker) string {
	name := fmt.Sprintf("tracking-%p", tracker)
	trackingDrivers.Store(name, tracker)
	sql.Register(name, &trackingDriver{name: name})
	return name
}

type trackingDriver struct{ name string }

func (d *trackingDriver) Open(string) (driver.Conn, error) {
	tracker, _ := trackingDrivers.Load(d.name)
	return &trackingConn{tracker: tracker.(*stmtTracker)}, nil
}

type trackingConn struct{ tracker *stmtTracker }

func (c *trackingConn) Prepare(query string) (driver.Stmt, error) {
	c.tracker.addPrepared(query)
	return &trackingStmt{tracker: c.tracker, query: query}, nil
}
func (c *trackingConn) Close() error { return nil }
func (c *trackingConn) Begin() (driver.Tx, error) {
	return nil, fmt.Errorf("transactions not supported")
}

type trackingStmt struct {
	tracker *stmtTracker
	query   string
	once    sync.Once
}

func (s *trackingStmt) Close() error {
	s.once.Do(func() { s.tracker.addClosed(s.query) })
	return nil
}
func (s *trackingStmt) NumInput() int { return -1 }
func (s *trackingStmt) Exec([]driver.Value) (driver.Result, error) {
	return driver.RowsAffected(1), nil
}
func (s *trackingStmt) Query([]driver.Value) (driver.Rows, error) { return &trackingRows{}, nil }

type trackingRows struct{}

func (*trackingRows) Columns() []string         { return []string{"x"} }
func (*trackingRows) Close() error              { return nil }
func (*trackingRows) Next([]driver.Value) error { return io.EOF }

// sleepUntil sleeps until deadline, busy-waiting the final stretch for
// millisecond-level accuracy.
func sleepUntil(deadline time.Time) {
	for {
		remaining := time.Until(deadline)
		if remaining <= 0 {
			return
		}
		if remaining > 25*time.Millisecond {
			time.Sleep(remaining - 20*time.Millisecond)
		}
	}
}
