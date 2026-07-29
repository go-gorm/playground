package main

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"io"
	"strings"
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

// TestEvictedStatementClosedBeforeExecution reproduces a use-after-close
// race in GORM's prepared statement cache (prepare_stmt.go +
// internal/stmt_store + internal/lru):
//
// PreparedStmtDB.ExecContext/QueryContext look up the cached *Stmt, release
// all cache locks, and only then call stmt.ExecContext. In between, the
// statement can be evicted — by TTL cleanup or by size-based LRU eviction —
// and stmt_store's onEvict callback closes it asynchronously. The in-flight
// execution then fails with "sql: statement is closed".
//
// That error is not driver.ErrBadConn, so the cache entry is not removed
// (prepare_stmt.go only deletes on ErrBadConn) and the request is not
// retried: a hot, correctly-configured cache can fail live requests. Note
// that cache hits do not extend an entry's TTL, so frequently-used
// statements are evicted just as easily as idle ones.
//
// This test deterministically reproduces the exact victim state of that
// race: it obtains a cached statement through the same store lookup that
// prepare() performs, lets eviction close it underneath (as happens in the
// lookup-to-execution window), and then executes it — the failure any
// concurrent request hits when the race is lost.
func TestEvictedStatementClosedBeforeExecution(t *testing.T) {
	tracker := &stmtTracker{preparedAt: map[string][]time.Time{}, closedAt: map[string][]time.Time{}}
	sqlDB, err := sql.Open(registerTrackingDriver(tracker), "")
	if err != nil {
		t.Fatal(err)
	}
	defer sqlDB.Close()
	sqlDB.SetMaxOpenConns(1)

	const q1, q2 = "SELECT 1", "SELECT 2"

	// maxSize=1 so adding a second statement evicts the first; a long TTL
	// keeps the background cleanup goroutine out of the picture.
	pdb := gorm.NewPreparedStmtDB(sqlDB, 1, time.Hour)
	defer pdb.Close()

	exec := func(query string) {
		t.Helper()
		if _, err := pdb.ExecContext(context.Background(), query); err != nil {
			t.Fatalf("exec %q: %v", query, err)
		}
	}

	exec(q1)

	// The same store lookup that PreparedStmtDB.prepare() performs right
	// before handing the statement to ExecContext.
	stmt, ok := pdb.Stmts.Get(q1)
	if !ok {
		t.Fatalf("expected %q to be cached", q1)
	}

	// Evicts q1's statement; stmt_store's onEvict closes it asynchronously,
	// exactly as TTL cleanup or LRU eviction does while another goroutine is
	// between prepare() and stmt.ExecContext.
	exec(q2)

	deadline := time.Now().Add(5 * time.Second)
	for {
		_, closed := tracker.counts(q1)
		if closed == 1 {
			break
		}
		if time.Now().After(deadline) {
			t.Fatal("evicted statement was not closed within 5s")
		}
		time.Sleep(time.Millisecond)
	}

	// The request that "lost the race": it holds a statement the cache has
	// already closed.
	_, err = stmt.ExecContext(context.Background())
	if err == nil {
		// The statement survived eviction; the race is fixed.
		return
	}
	if !strings.Contains(err.Error(), "statement is closed") {
		t.Fatalf("unexpected error executing evicted statement: %v", err)
	}
	if errors.Is(err, driver.ErrBadConn) {
		t.Fatalf("closed-statement error should not be reported as driver.ErrBadConn: %v", err)
	}

	t.Fatalf("cached statement was closed by eviction before execution could use it; "+
		"the resulting %q is not driver.ErrBadConn, so PreparedStmtDB neither re-prepares nor retries and the request fails", err)
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
