package main

import (
	"errors"
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	// error is here
	// when use find the error is nill
	if err := DB.Find(&result, "id = ?", 0).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("Record not found")
	}

	if err := DB.First(&result, "id = ?", 0).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("Record not found")
	}
}
