package main

import (
	"testing"

	"github.com/pioz/faker"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: mysql, postgres

func TestGORM(t *testing.T) {
	faker.SetSeed(1) // Get determinist output
	products := make([]Product, 100)
	faker.Build(&products)
	err := DB.Create(&products).Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
