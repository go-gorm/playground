package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	users := []User{
		{Name: "a", Age: 1},
		{Name: "b", Age: 2},
		{Name: "c", Age: 3},
	}

	DB.Create(&users)

	var result int64
	if err := DB.Model(&User{}).Not(DB.Where("name = 'a'").Or("name = 'b'")).Count(&result).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	if result != 1 {
		t.Errorf("Failed, expect: 1, got: %v", result)
	}
}

