package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	account := Account{Number: "77"}
	DB.Create(&account)

	user := User{Name: "jinzhu", Account: account}

	DB.Create(&user)

	var result, result2, result3, result4 User
	if err := DB.Preload("Account").First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	if result.Account.ID != account.ID {
		t.Fatalf("Failed using preload, Account should be valid")
	}
	if err := DB.Joins("Account").First(&result2, user.ID).Error; err != nil {
		t.Fatalf("Failed, got error: %v", err)
	}
	if result2.Account.ID != account.ID {
		t.Fatalf("Failed using joins, account should be valid")
	}

	// Delete the account
	DB.Delete(&account)
	if err := DB.Preload("Account").First(&result3, user.ID).Error; err != nil {
		t.Fatalf("Failed, got error: %v", err)
	}
	if result3.Account.ID != 0 {
		t.Errorf("Failed using preload, Account should be 0 got: %v", result3)
	}

	if err := DB.Joins("Account").First(&result4, user.ID).Error; err != nil {
		t.Fatalf("Failed, got error: %v", err)
	}
	if result4.Account.ID != 0 {
		t.Errorf("Failed using joins, account should be 0, got: %v", result4)
	}

}
