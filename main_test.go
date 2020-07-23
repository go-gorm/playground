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

	var result string
	if err := DB.Table("users").Select("id").First(&result).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
