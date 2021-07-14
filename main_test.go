package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}
	pet := Pet{UserID: &user.ID, Name: "jinzhu"}

	DB.Create(&user)
	DB.Create(&pet)

	// This is fine
	if err := DB.Model(&User{}).Distinct("users.*").Joins("JOIN pets ON users.id = pets.user_id").Find(&user).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	// This is also fine
	var count int64
	if err := DB.Model(&User{}).Distinct("users.id").Joins("JOIN pets ON users.id = pets.user_id").Count(&count).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	// This is not
	if err := DB.Model(&User{}).Distinct("users.*").Joins("JOIN pets ON users.id = pets.user_id").Count(&count).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
