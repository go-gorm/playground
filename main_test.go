package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	users := []*User{{Name: "jinzhu"}}

	if err := DB.Create(users).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
