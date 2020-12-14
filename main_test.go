package main

import (
	"testing"

	gorm "gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	zero := gorm.DeletedAt{}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if result.Model.DeletedAt != zero {
		t.Errorf("Failed, DeletedAt should be empty")
	}

	if err := DB.Delete(&result).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if result.Model.DeletedAt == zero {
		t.Errorf("Failed, DeletedAt is empty")
	}
}
