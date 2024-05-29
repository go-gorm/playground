package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu", State: NewCustomType("active")}

	DB.Create(&user)

	var result []User
	if err := DB.Find(&result).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if len(result) == 0 {
		t.Errorf("Failed, got empty result")
	}

	if result[0].State.Machine == nil {
		t.Errorf("Failed, expected: active")
	}
}
