package main

import (
	"testing"

	"gorm.io/datatypes"
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

func TestSaveJSONNil(t *testing.T) {
	thing := ThingWithJSON{Data: nil}

	err := DB.Save(&thing).Error
	if err != nil {
		t.Errorf("Failed to save thing with json, got error: %w", err)
	}
}

func TestSaveJSON(t *testing.T) {
	thing := ThingWithJSON{Data: datatypes.JSON(`{"foo": "bar"}`)}

	err := DB.Save(&thing).Error
	if err != nil {
		t.Errorf("Failed to save thing with json, got error: %w", err)
	}
}
