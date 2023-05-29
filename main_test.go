package main

import (
	"testing"
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

	err := DB.Model(&result).Association("Languages").Append(&Language{
		Code: "en-us",
		Name: "English",
	})
	if err != nil {
		t.Errorf("Association Append Failed, got error: %v", err)
	}

	user2 := User{Name: "diligiant"}

	DB.Create(&user2)

	var result2 User
	if err := DB.First(&result2, user2.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	err = DB.Model(&result2).Association("Languages").Append(&Language{
		Code: "en-us",
		Name: "English",
	})
	if err != nil {
		t.Errorf("Association Append Failed, got error: %v", err)
	}

	var languages []Language
	if err := DB.Find(&languages).Error; err != nil {
		t.Errorf("Finding languages Failed, got error: %v", err)
	}
}
