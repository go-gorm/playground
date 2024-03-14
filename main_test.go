package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	users := []User{}
	DB.Create(&User{
		Name: "jinzhu",
	})

	DB.Where("status = ?", Normal).Find(&users)
	DB.Where("status = ?", 1).Find(&users)
}
