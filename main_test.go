package main

import (
	"gorm.io/playground/diff"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	if err := DB.AutoMigrate(&User{}); err != nil {
		t.Fatal(err)
	}

	if err := DB.AutoMigrate(&diff.User{}); err != nil {
		t.Fatal(err)
	}
}
