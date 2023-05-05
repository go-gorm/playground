package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	go func() {
		cms := make([]Company, 0)
		DB.Find(&cms)
	}()
	users := make([]User, 0)
	DB.Preload("Toys").Find(&users)
}
