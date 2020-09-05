package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: v1.20.0
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	type Product struct {
		Value float64 `gorm:"default:1"`
	}

	for i := 0; i < 2; i++ {
		if err := DB.AutoMigrate(&Product{}); err != nil {
			t.Errorf("Failed, got error: %v", err)
		}
	}
}
