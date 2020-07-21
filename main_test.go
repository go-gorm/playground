package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	comp := Company{
		Name: "Test",
	}
	DB.Create(&comp)
	data := UserWork{}
	DB.AutoMigrate(&data)
}
