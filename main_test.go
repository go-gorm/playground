package main

import (
	"context"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DefaultNullDemo struct {
	ID  uint   `gorm:"primarykey"`
	Key string `gorm:"size:255;default:NULL"`
}

type Counter struct {
	logger.Interface
	count int64
}

func (a *Counter) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	sql, _ := fc()
	if strings.HasPrefix(strings.ToUpper(strings.TrimSpace(sql)), "ALTER TABLE") {
		atomic.AddInt64(&a.count, 1)
	}
	a.Interface.Trace(ctx, begin, fc, err)
}

func TestPGDefaultNullAlterColumn(t *testing.T) {

	// 首次：CreateTable
	if err := DB.AutoMigrate(&DefaultNullDemo{}); err != nil {
		t.Fatal(err)
	}

	// 第二次：装上计数 logger，期望 0 条 ALTER
	counter := &Counter{Interface: DB.Logger}
	session := DB.Session(&gorm.Session{Logger: counter})
	if err := session.AutoMigrate(&DefaultNullDemo{}); err != nil {
		t.Fatal(err)
	}
	if err := DB.AutoMigrate(&DefaultNullDemo{}); err != nil {
		t.Fatal(err)
	}

	if counter.count > 0 {
		t.Fatalf("expected 0 ALTER statements on second AutoMigrate, got %d", counter.count)
	} else {
		t.Log("ok")
	}
}
