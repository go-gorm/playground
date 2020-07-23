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
	t.Run("take", func(t *testing.T) {
		if err := DB.Table("users").Select("name").Take(&result).Error; err != nil {
			t.Errorf("Failed, result: %v, got error: %v", result, err)
		}
	})
	t.Run("First", func(t *testing.T) {
		if err := DB.Table("users").Select("name").First(&result).Error; err != nil {
			t.Errorf("Failed, got error: %v", err)
		}
	})
}
