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

	user.Name = "test"

	if err := DB.Save(user).Error; err != nil {
		t.Errorf("Failed to save 1, got error: %v", err)
	}

	if err := DB.Model(&User{}).Save(user).Error; err != nil {
		t.Errorf("Failed to save 3, got error: %v", err)
	}

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
