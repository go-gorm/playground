package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu", Pets: []*Pet{{Name: "123"}}}

	DB.Create(&user)

	user.Pets[0].Name = "1234"
	DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user)

	var result User
	if err := DB.Preload("Pets").First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if result.Pets[0].Name != "1234" {
		t.Errorf("Failed, exp 1234 act: %s", result.Pets[0].Name)
	}
}
