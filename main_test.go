package main

import (
	"fmt"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	var account = NewArticle{Title: "test1"}
	if err := DB.Create(&account).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
		return
	}

	if account.ID == 0 {
		t.Errorf("Failed, got error: %s", "id is zero")
	}

	fmt.Printf("ID: %d \n", account.ID)

	var accounts = []NewArticle{{Title: "test2"}, {Title: "test3"}}
	if err := DB.Create(&accounts).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
		return
	}

	for _, row := range accounts {
		if row.ID == 0 {
			t.Errorf("Failed, got error: %s", "id is zero")
		}

		fmt.Printf("ID: %d \n", row.ID)
	}
}
