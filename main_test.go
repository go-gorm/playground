package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	company := Company{Name: "jinzhu-co", Address: "co-address"}

	DB.Create(&company)

	user := User{Name: "jinzhu", CompanyID: &company.ID}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if err := DB.Preload("Company").First(&result).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if result.Company.Name == "" {
		t.Errorf("Failed, got empty company")
	}

	if err := DB.Preload("Company", func(db *gorm.DB) *gorm.DB {
		return db.Select("*")
	}).First(&result).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if result.Company.Name == "" {
		t.Errorf("Failed with * select, got empty company")
	}

	if err := DB.Preload("Company", func(db *gorm.DB) *gorm.DB {
		return db.Select("name")
	}).First(&result).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if result.Company.Name == "" {
		t.Errorf("Failed with fields select, got empty company")
	}
}
