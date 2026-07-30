package main

import (
	"context"
	"database/sql/driver"
	"errors"
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

// TestPreparedStmtEvictionRace reproduces a use-after-close race in GORM's
// prepared statement cache.
//
// PreparedStmtDB.ExecContext/QueryContext look up the cached *Stmt, release
// every cache lock, and only then call stmt.ExecContext. In that window another
// caller can evict the statement — by LRU size here, by TTL cleanup in
// production — and stmt_store's onEvict closes it asynchronously. The in-flight
// execution then fails with "sql: statement is closed".
//
// That error is not driver.ErrBadConn, and prepare_stmt.go drops a cache entry
// only on ErrBadConn, so the statement is neither re-prepared nor retried: the
// request just fails. Cache hits do not extend an entry's TTL either, so hot
// statements are evicted as readily as idle ones.
//
// Two callers, two queries, cache size 1: each call evicts the other's
// statement while that other call sits between prepare() and Exec.
// NewPreparedStmtDB is what gorm.Open builds for
// gorm.Config{PrepareStmt: true, PrepareStmtMaxSize: 1}.
func TestPreparedStmtEvictionRace(t *testing.T) {
	sqlDB, err := DB.DB()
	if err != nil {
		t.Fatal(err)
	}

	pdb := gorm.NewPreparedStmtDB(sqlDB, 1, time.Hour)
	defer pdb.Close()

	var wg sync.WaitGroup
	failures := make(chan error, 2)
	for _, query := range []string{"SELECT 1", "SELECT 2"} {
		wg.Add(1)
		go func(query string) {
			defer wg.Done()
			for i := 0; i < 5000; i++ {
				if _, err := pdb.ExecContext(context.Background(), query); err != nil {
					failures <- err
					return
				}
			}
		}(query)
	}
	wg.Wait()
	close(failures)

	err, lost := <-failures
	if !lost {
		t.Log("no execution landed in the eviction window this run")
		return
	}
	if !strings.Contains(err.Error(), "statement is closed") {
		t.Fatalf("unexpected error from a cached statement: %v", err)
	}
	if errors.Is(err, driver.ErrBadConn) {
		t.Fatalf("expected the closed-statement error not to be driver.ErrBadConn, got: %v", err)
	}
	t.Fatalf("cached statement was closed by eviction between lookup and execution: %v; "+
		"this is not driver.ErrBadConn, so PreparedStmtDB neither re-prepares nor retries and the request fails", err)
}
