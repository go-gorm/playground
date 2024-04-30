package main

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/logger"
	"testing"
	"time"
)

var actualSql string

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {

	DB.DryRun = true
	DB.Logger = &gormLogger{}
	var result User

	// work well case:
	expectSql := "SELECT * FROM `users` WHERE (`field1` = 1 OR `field2` = 2) AND `field3` = 3 AND `users`.`deleted_at` IS NULL"
	if err := DB.
		Where("`field1` = 1 OR `field2` = 2").
		Where("`field3` = 3").
		Find(&result).Error; err != nil {
		t.Error(err)
	}
	assert.Equal(t, expectSql, actualSql)

	// issue case:
	expectSql = "SELECT * FROM `users` WHERE (`field1` = 1 || `field2` = 2) AND `field3` = 3 AND `users`.`deleted_at` IS NULL"
	if err := DB.
		Where("`field1` = 1 || `field2` = 2").
		Where("`field3` = 3").
		Find(&result).Error; err != nil {
		t.Error(err)
	}
	assert.Equal(t, expectSql, actualSql)

}

type gormLogger struct{}

func (g *gormLogger) LogMode(_ logger.LogLevel) logger.Interface {
	newLogger := *g
	return &newLogger
}

func (g *gormLogger) Info(context.Context, string, ...any) {}

func (g *gormLogger) Warn(context.Context, string, ...any) {}

func (g *gormLogger) Error(context.Context, string, ...any) {}

func (g *gormLogger) Trace(c context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	actualSql = sql
}
