package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	if err := DB.AutoMigrate(&User{}); err!=nil {
		t.Errorf("Failed, migrate error: %v", err)
	}

	// do twice will failed
	if err := DB.AutoMigrate(&User{}); err!=nil {
		t.Errorf("Failed, migrate error: %v", err)
	}
}
