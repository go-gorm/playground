package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	u := user{Name: "jinzhu"}

	DB.Create(&u)

	var result user
	if err := DB.First(&result, u.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
