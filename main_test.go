package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu", SomeCustomString: "test"}

	DB.Create(&user)

	// the following statement works
	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	if result.SomeCustomString != "test" {
		t.Errorf("Failed, expected: %v, got: %v", "test", result.SomeCustomString)
	}

	// the following statement works
	var strResult string
	if err := DB.Model(&User{}).Where("id = ?", user.ID).Select("some_custom_string").Find(&strResult).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	// the following statement does not work
	var customStrResult CustomString
	if err := DB.Model(&User{}).Where("id = ?", user.ID).Select("some_custom_string").Find(&customStrResult).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
