package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	if DB.Dialector.Name() == "sqlite" {
		DB.Exec("PRAGMA foreign_keys = ON")
	}
	RunMigrations()
}
