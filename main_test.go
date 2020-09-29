package main

import (
	"testing"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	config := "user=postgres password=Zps05..... dbname=version3 host=localhost port=5432 sslmode=disable"
	for {
		db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
		if err != nil {
			println(err.Error())
		}
		sqlDB, err := db.DB()
		_ = sqlDB.Close()
		time.Sleep(10 * time.Millisecond)
	}
}
