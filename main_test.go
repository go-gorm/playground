package main

import (
	"testing"
)



// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user1 := User{Name: "jinzhu", Age: 42}
	user2 := User{Name: "jinzhu", Age: 99}

	DB.Create(&user1)
	DB.Create(&user2)

	correctResults := []User{}
	if err := DB.Raw("SELECT DISTINCT(name), age FROM users").Scan(&correctResults).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	if len(results) != 1 {
		t.Errorf("Expected 1 result, got %d. %v", len(correctResults), correctResults)
	}

	actualResults := []User{}
	if err := DB.Model(&User{}).Distinct("name").Select("name", "age").Scan(&actualResults).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	if len(actualResults) != 1 {
		t.Errorf("Expected 1 result, got %d. %v", len(actualResults), actualResults)
	}
}
