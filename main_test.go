package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: mysql

func TestGORM(t *testing.T) {
	companyID := 100
	user := User{
		Name:      "user_A",
		CompanyID: &companyID,
		Languages: []Language{
			{Name: "language_A", Code: 10, CompanyID: &companyID},
		},
	}

	DB.Create(&user)
	var result1 User
	if err := DB.Preload("Languages").First(&result1, 1).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	DB.Create(&User{Name: "user_B"})
	var result2 User
	if err := DB.Preload("Languages").First(&result2, 2).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
