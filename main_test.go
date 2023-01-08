package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var results []User
	if err := DB.Where(&User{
		Active: false,
	}, "Active").Find(&results, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	// This must have 2 results since the first resource was created when the column did not exist yet
	// I therefore assume that it should have the default value when using AutoMigrate.
	// The second record is created when the column exits and therefore has the default value set.
	assert.Len(t, results, 2)
}
