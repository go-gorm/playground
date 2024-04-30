package main

import (
	"errors"
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			t.Errorf("Failed, got error: %v", err)
		}
	}
}

func TestReMigrate(t *testing.T) {
	// re-migrate existing table
	u := &User{}
	if !DB.Migrator().HasTable(u) {
		if err := DB.AutoMigrate(u); err != nil {
			t.Fatalf("Failed to auto migrate, but got error %v\n", err)
		}
	}
	if err := DB.AutoMigrate(u); err != nil {
		t.Fatalf("Failed to auto migrate, but got error %v\n", err)
	}
}
