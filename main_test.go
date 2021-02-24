package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{
		Name: "jinzhu",
		Age:  32,
		Account: Account{
			Number: "123",
		},
	}

	DB.Create(&user)
	db := DB.Where("FakeName = 1")

	tx := db.Session(&gorm.Session{NewDB: true}).Begin()

	var acc Account
	if err := tx.Where("number = ?", "345").Find(&acc).Error; err != nil {
		t.Errorf("error finding account: %v", err)
		tx.Rollback()
		return
	}
	tx.Commit()
	return

}
