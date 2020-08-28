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

	var result1 User
	if err := DB.Table("user").Where("name = ?", "tpp").Find(&result1).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	
	var result2 User
	if err := DB.Table("user").Where("name = ?", "tpp").First(&result2).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
