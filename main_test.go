package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Session(&gorm.Session{}).Create(&user)

	var result User

	if err := DB.First(&result, user.ID).Update("name", "jinzhu 2").Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
