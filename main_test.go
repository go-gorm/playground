package main

import (
	"testing"
	
	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	account := Account{}

	user := User{
		Name: "jinzhu",
		Account: account,
	}

	DB.Create(&user)

	var result User
	if err := DB.Session(&gorm.Session{SkipHooks: true}).Preload("Account").First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	
	if result.Account.Number == "123" {
		t.Errorf("SkipHooks did not skip a Preloaded relation's hook")
	}
}
