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

	if err := DB.Select("u.id, u.*").Table("users as u").First(&result, user.ID).
		Find(&result).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if err := DB.Select("u.*").Table("users as u").First(&result, user.ID).
		Find(&result).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
