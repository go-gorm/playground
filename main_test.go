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

	var c int64

	if err := DB.Table("users").Where("name = ?", "jinzhu").Select("name").Count(&c); err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if err := DB.Table("users").Where("name = ?", "jinzhu").Count(&c); err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if err := DB.Table("users").Count(&c); err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}