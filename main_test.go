package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type TableA struct {
	Name string `gorm:"size:100;-:migration"`
}

func TestGORM(t *testing.T) {
	if err := DB.AutoMigrate(&TableA{}); err != nil {
		t.Errorf("failed migrate: %v", err)
		return
	}
}
