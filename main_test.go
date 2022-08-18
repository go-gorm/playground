package main

import (
	"testing"

	"github.com/google/uuid"
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
}

func TestScanToArray(t *testing.T) {
	var someUUID *uuid.UUID
	if err := DB.Table("users").Select("some_uuid").Scan(&someUUID).Error; err != nil {
		t.Errorf("Failed, go error: %v", err)
	}
}
