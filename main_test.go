package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu", Active: true}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if err := DB.Where("id = ?", user.ID).
		Assign(&User{Name: "new_name", Active: false}).
		FirstOrCreate(&User{}).Error; err != nil {
		t.Errorf("Failed to update user")
	}

	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if result.Name != "new_name" {
		t.Errorf("name did not updated")
	}

	if result.Active {
		t.Errorf("active did not updated")
	}
}
