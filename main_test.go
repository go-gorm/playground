package main

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

// recordingPool records the context gorm passes to BeginTx.
type recordingPool struct {
	*sql.DB
	beginCtx context.Context
}

func (p *recordingPool) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	p.beginCtx = ctx
	return p.DB.BeginTx(ctx, opts)
}

func TestGORM(t *testing.T) {
	sqlDB, err := DB.DB()
	if err != nil {
		t.Fatalf("DB: %v", err)
	}

	db := DB.Session(&gorm.Session{NewDB: true})
	db.Config.DefaultTransactionTimeout = time.Hour
	pool := &recordingPool{DB: sqlDB}
	db.Statement.ConnPool = pool

	tx := db.Begin()
	if tx.Error != nil {
		t.Fatalf("Begin: %v", tx.Error)
	}
	if err := tx.Commit().Error; err != nil {
		t.Fatalf("Commit: %v", err)
	}

	if _, ok := pool.beginCtx.Deadline(); !ok {
		t.Fatal("Begin did not apply DefaultTransactionTimeout")
	}
	if pool.beginCtx.Err() == nil {
		t.Error("the DefaultTransactionTimeout context is still live after Commit; its timer runs for the full timeout")
	}
}
