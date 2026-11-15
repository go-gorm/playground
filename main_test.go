package main

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"gorm.io/gorm"
	"gorm.io/playground/models"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Simulate a client disconnect during a write: cancel the request context
	// after gorm has opened its default transaction, and wait until database/sql
	// has rolled that transaction back on its own.
	err := DB.Callback().Create().Before("gorm:create").Register("cancel_ctx", func(tx *gorm.DB) {
		sqlTx := tx.Statement.ConnPool.(*sql.Tx)
		cancel()
		deadline := time.Now().Add(5 * time.Second)
		for time.Now().Before(deadline) {
			if _, err := sqlTx.ExecContext(context.Background(), "SELECT 1"); errors.Is(err, sql.ErrTxDone) {
				return
			}
		}
	})
	if err != nil {
		t.Fatal(err)
	}

	user := models.User{Name: "jinzhu"}
	err = DB.WithContext(ctx).Create(&user).Error

	if !errors.Is(err, context.Canceled) {
		t.Errorf("expected errors.Is(err, context.Canceled), got: %v", err)
	}
}
