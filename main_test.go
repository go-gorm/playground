package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlserver

type Parent struct {
	gorm.Model
	Name string `gorm:"uniqueIndex;size:200"`
}

func TestGORM(t *testing.T) {
	const parentName = "John"
	user1 := Parent{Name: parentName}
	user2 := Parent{Name: parentName}

	_ = DB.AutoMigrate(&Parent{})

	err := DB.Create(&user1).Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	err = DB.Create(&user2).Error
	if err == nil {
		t.Errorf("Failed, must raise duplicate key error, got: %v", err)
	}
}
