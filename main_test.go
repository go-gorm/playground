package main

import (
	"gorm.io/gorm"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestPreload(t *testing.T) {
	//1.Add User
	user := User{Name: "jinzhu"}
	DB.Create(&user)
	//2.Add 2 Accounts
	accounts := []Account{{UserID: user.ID, Number: "1"}, {UserID: user.ID, Number: "2"}}
	DB.Create(accounts)

	var result User
	err := DB.Preload("Account").First(&result).Error
	if err != nil {
		t.Fail()
	}
	if len(result.Account) != 2 {
		t.Fail()
	}
}

func TestPreloadWithSelect(t *testing.T) {
	//1.Add User
	user := User{Name: "jinzhu"}
	DB.Create(&user)
	//2.Add 2 Accounts
	accounts := []Account{{UserID: user.ID, Number: "1"}, {UserID: user.ID, Number: "2"}}
	DB.Create(accounts)

	var result User
	err := DB.Preload("Account", func(db *gorm.DB) *gorm.DB {
		return db.Select("number")
	}).First(&result).Error
	if err != nil {
		t.Fail()
	}
	if len(result.Account) != 2 {
		t.Fail()
	}
}
