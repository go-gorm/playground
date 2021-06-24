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

	DB.Create(&user)

	var (
		count int64
		query = DB.Model(&User{}).Order("users.id ASC")
		session = query.Session(&gorm.Session{})
	)
	if err := query.Count(&count).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	if err := session.Count(&count).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
