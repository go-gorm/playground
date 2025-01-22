package main

import (
	"testing"
	"time"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{
		Name: "jinzhu",
	}
	DB.Create(&user)

	bday := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	user2 := User{
		Name:     "sam-wbcx",
		Birthday: &bday,
	}
	DB.Create(&user2)

	var results []User
	if err := DB.
		Where(`
			birthday IS NULL
			OR birthday < CURRENT_TIMESTAMP
		`).
		Where("name = ?", user2.Name).
		Find(&results).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if len(results) != 1 {
		t.Errorf("Failed, returned %d results", len(results))
	}
}
