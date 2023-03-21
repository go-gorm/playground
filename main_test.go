package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	company := Company{Name: "jinzhu"}

	DB.Create(&company)

	var result Company

	DB.First(&result, company.ID)

	if err := DB.Save(&result).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
