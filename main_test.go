package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: v1.25.5
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	users := []User{
		{Name: "a", Age: 1},
		{Name: "a", Age: 2},
		{Name: "b", Age: 1},
		{Name: "b", Age: 2},
	}

	DB.Create(&users)

	var result int64
	if err := DB.Model(&User{}).Not(DB.Where("name = 'a'").Where("age = 1")).Count(&result).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	if result != 3 {
		t.Errorf("Failed, expect: 3, got: %v", result)
	}
}
