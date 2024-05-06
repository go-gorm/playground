package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu", Languages: []Language{{Code: "ZH", Name: "Chinese"}, {Code: "EN", Name: "English"}}}

	DB.Create(&Language{})
	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.Email).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
