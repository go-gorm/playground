package main

import (
	"context"
	"database/sql"
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

// TestExpiredPreparedStatementLeak reproduces a prepared statement leak in
// GORM's statement cache.
//
// internal/lru.Get reports a miss for an expired entry but leaves that entry in
// the cache. Re-running the query before the cleanup goroutine collects it takes
// lru.Add's existing-item branch, which overwrites ent.Value without calling
// onEvict. onEvict (internal/stmt_store.New) is the only place cached statements
// are closed, so the replaced *sql.Stmt is never closed and nothing references it
// any more: one leaked statement per TTL cycle per hot query, released only when
// the whole *sql.DB is closed.
//
// NewPreparedStmtDB below is what gorm.Open builds for
// gorm.Config{PrepareStmt: true, PrepareStmtTTL: ttl}. The counting driver in
// counting_driver_test.go is only there to observe Prepare/Close, which
// database/sql does not expose.
func TestExpiredPreparedStatementLeak(t *testing.T) {
	const ttl = 10 * time.Second
	const tick = ttl / 100 // lru cleanup granularity: one expiration bucket per tick

	sqlDB, err := sql.Open(countingDriverName, "")
	if err != nil {
		t.Fatal(err)
	}
	defer sqlDB.Close()

	start := time.Now()
	pdb := gorm.NewPreparedStmtDB(sqlDB, 100, ttl)
	defer pdb.Close()

	exec := func(query string) {
		t.Helper()
		if _, err := pdb.ExecContext(context.Background(), query); err != nil {
			t.Fatalf("exec %q: %v", query, err)
		}
	}

	// Put both statements in the same expiration bucket, a quarter tick in so
	// neither lands in a neighbouring one.
	time.Sleep(time.Until(start.Add(time.Since(start)/tick*tick + tick + tick/4)))

	t0 := time.Now()
	exec("SELECT 1") // expires at t0+ttl, in bucket B
	time.Sleep(tick / 2)
	exec("SELECT 2") // same bucket, expires half a tick later, so B's cleanup waits that long

	// Re-run the first query once it has expired but before bucket B is collected.
	time.Sleep(time.Until(t0.Add(ttl + 20*time.Millisecond)))
	exec("SELECT 1")

	time.Sleep(3 * tick) // let bucket B's cleanup pass finish
	prepared, closed := stmtsPrepared.Load(), stmtsClosed.Load()
	t.Logf("prepared=%d closed=%d", prepared, closed)

	if prepared == 3 && closed == 1 {
		t.Fatal("leaked prepared statement: the expired cache entry was overwritten without Close, " +
			"so the replaced *sql.Stmt stays open on the *sql.DB with nothing left referencing it")
	}
}
