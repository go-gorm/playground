package main

import (
	"strings"
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	ss := DB.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Limit(1)
	}).Session(&gorm.Session{
		DryRun: true,
	})
	ss.Migrator()
	stm := ss.Find(&User{}).Statement
	sql := stm.SQL.String()
	if !strings.Contains(sql, "LIMIT") {
		t.Errorf("scope was not apply")
	}
}
