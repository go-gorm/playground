package main

import (
	"testing"
)

type User struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	CompanyID  int

	Company    *Company `gorm:"embedded;embeddedPrefix:com"`
}

type Company struct {
	ID        int `gorm:"primaryKey"`
	Name      string
}

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu", CompanyID: 1}
	if err := DB.Create(&user).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
