package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var results []User
	if err := DB.Find(&results, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	// This must have 1 result since we have only created one resource.
	assert.Len(t, results, 1)
}
