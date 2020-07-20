package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	var account = Account{Number: "test1"}
	if err := DB.Create(&account).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
		return
	}

	if account.ID == 0 {
		t.Errorf("Failed, got error: %s", "id is zero")
	}

	var accounts = []Account{{Number: "test2"}, {Number: "test3"}}
	if err := DB.Create(&accounts).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
		return
	}

	for _, row := range accounts {
		if row.ID == 0 {
			t.Errorf("Failed, got error: %s", "id is zero")
		}
	}
}
