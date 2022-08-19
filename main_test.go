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

	expectedUUID := uuid.New()
	user := User{Name: "foo", SomeUUID: expectedUUID}
	DB.Create(&user)

	if err := DB.Table("users").Select("some_uuid").Where("name = ?", "foo").Scan(&someUUID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if someUUID == nil {
		t.Error("someUUID is nil")
	}
	if *someUUID != expectedUUID {
		t.Errorf("someUUID (%v) != expectedUUID (%v)", *someUUID, expectedUUID)
	}
}
