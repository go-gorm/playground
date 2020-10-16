package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	if err := DB.AutoMigrate(&SomeModel{}); err != nil {
		t.Errorf("Failed SomeModel, got error: %v", err)
	}
	if err := DB.AutoMigrate(&SomeAttModel{}); err != nil {
		t.Errorf("Failed SomeAttModel, got error: %v", err)
	}
}
