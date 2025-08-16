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
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}

func TestCountAndTables(t *testing.T) {
	var cnt int64
	if err := DB.Table("sessions AS s").
		Table("users").Count(&cnt).Error; err != nil {
		t.Errorf("got error on calling Count(): %v", err)
	}
}
