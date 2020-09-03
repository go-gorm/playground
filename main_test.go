package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	rawDB, rawErr := DB.DB()
	if rawErr != nil {
		t.Errorf("Failed, got error: %v", rawErr)
	}
	rawDB.Close()

	DB.Transaction(func(tx *gorm.DB) error {
		return nil
	})
}
