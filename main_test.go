package main

import (
	"gorm.io/gorm"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	if err := DB.Transaction(func(tx *gorm.DB) error {
		return tx.AutoMigrate(&User{})
	}); err != nil {
		t.Error(err)
	}

	if err := DB.Transaction(func(tx *gorm.DB) error {
		return tx.AutoMigrate(&User{})
	}); err != nil {
		t.Error(err)
	}
}
