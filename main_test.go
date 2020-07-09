package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type Test struct {
	Array []string
}


func TestGORM(t *testing.T) {
	err := DB.AutoMigrate(&Test{})
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
