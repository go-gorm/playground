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
}

func TestCount(t *testing.T) {
	query := DB.Model(UserWithPetCount{}).Select("users.*, COUNT(pets.id) AS pet_count").
		Joins("LEFT JOIN pets ON pets.user_id = users.id").
		Group("users.id")

	var count int64 

	if err := query.Count(&count).Error; err != nil {
		t.Errorf("Failed count query, got error: %v", err)
	}
	
	var users []UserWithPetCount

	if err := query.Limit(5).Find(&users).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
