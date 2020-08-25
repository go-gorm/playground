package main

import (
	"fmt"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}
	user.Account = Account{Number: "10000"}
	DB.Create(&user)

	var result User
	if err := DB.Preload("Account").First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	fmt.Print(result)

	result.Account.Number = "10001"
	DB.Save(&result)

	if err := DB.Preload("Account").First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	fmt.Print(result)
}
