package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	var user User
	user = User{Name: "noopple"}
	DB.Create(&user)
	var pet Pet
	pet = Pet{
		Name: "Marsel",
		UserID: &user.ID,
	}
	DB.Create(&pet)

	pet = Pet{
		Name: "Gendalf",
		UserID: &user.ID,
	}
	DB.Create(&pet)

	var userID uint

	err := DB.Model(&User{}).
		Select("users.id").
		Joins("Pets").
		Where("LENGTH(\"Pets\".name) > 3").
		Group("users.id").
		Having("COUNT(1) > 1").
		Scan(&userID).Error

	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
