package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: postgres

func TestGORM(t *testing.T) {
	DB.AutoMigrate(&Account{})
	user := Account{}
	user.ID = "jinzhu"

	if err := DB.Create(&user).Error; err != nil {
		t.Errorf("Could not create account")
	}

	var result Account
	if err := DB.Preload("Created").First(&result, user.PK).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if result.Created == nil {
		t.Errorf("Created not loaded")
	}
}
