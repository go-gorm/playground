package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type Crash struct {
	WithDefault  string `gorm:"default:DEFAULT1"`
	WithDefault2 string `gorm:"default:DEFAULT2"`
}

func TestGORM(t *testing.T) {
	DB.Migrator().DropTable(Crash{})
	DB.Migrator().CreateTable(Crash{})
	crashIt := Crash{}

	DB.Create(&crashIt)

	DB.AutoMigrate(Crash{})
}
