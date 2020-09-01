package main

import (
	"context"
	"gorm.io/gorm/logger"
	"sync/atomic"
	"testing"
	"time"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)
	dLogger := &dummyLogger{}
	DB.Logger = dLogger

	DB.Find(&User{})
	// when use model query
	if dLogger.rowEffected != 1 {
		t.Error("Wrong dummy logger setup")
		return
	}

	// same query as above
	rows, err := DB.Raw("SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL").Rows()
	if err != nil {
		t.Error("Error: ", err)
		return
	}

	// check return rows
	countRow := 0
	for rows.Next() {
		countRow++
	}
	if countRow != 1 {
		t.Error("Wrong setup")
		return
	}

	// expect logged returned rows
	if dLogger.rowEffected != 1 {
		t.Error("Expect logging for 1 returned row")
	}
}

type dummyLogger struct {
	rowEffected int64 // record number of effected rows (in tracing logger)
}

func (l *dummyLogger) LogMode(logger.LogLevel) logger.Interface {
	return l
}
func (l *dummyLogger) Info(context.Context, string, ...interface{}) {
}

func (l *dummyLogger) Warn(context.Context, string, ...interface{}) {

}
func (l *dummyLogger) Error(context.Context, string, ...interface{}) {

}
func (l *dummyLogger) Trace(_ context.Context, _ time.Time, fc func() (string, int64), _ error) {
	_, rows := fc()
	atomic.StoreInt64(&l.rowEffected, rows)
}
