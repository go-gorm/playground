package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	var total int64
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	if err := DB.
		Model(&user).
		Scopes(func(db *gorm.DB) *gorm.DB { return db.Select("*, 1 as a") }).
		Count(&total).
		Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
