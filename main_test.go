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
	t.Run("scan", func(t *testing.T) {
		if err := DB.Table("users").Select("name").Scan(&result).Error; err != nil || result == "" {
			t.Errorf("Failed, result: %v, got error: %v", result, err)
		}
	})
}
