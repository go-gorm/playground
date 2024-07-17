package main

import (
	"testing"
	"time"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	for i := 0; i < 1; i++ {
		go func() {
			var result User
			DB.First(&result)
		}()
	}

	for i := 0; i < 1; i++ {
		go func() {
			var result Account
			DB.First(&result)
		}()
	}

	time.Sleep(time.Second)
}
