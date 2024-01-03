package main

import (
	"gorm.io/gorm"
	"strings"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "     jinzhu     "}

	err := DB.Create(&user).Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	// Works fine with create
	if result.Name != "jinzhu" {
		t.Errorf("Failed, got name: %v. Expected: %v", result.Name, "jinzhu")
	}

	// Add a name with non-UTF8 characters so that the update will error if characters are not removed
	userBrokenName := User{Name: "     tux     "}
	if err := DB.Model(&user).Updates(userBrokenName).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	// Does not work with update
	if result.Name != "tux" {
		t.Errorf("Failed, got name: %v. Expected: tux", result.Name)
	}
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	cleanName := strings.TrimSpace(u.Name) // remove trailing whitespace
	u.Name = cleanName
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	cleanName := strings.TrimSpace(u.Name) // remove trailing whitespace
	u.Name = cleanName
	return nil
}
