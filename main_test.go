package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	type exampleWithComputedCols struct {
		gorm.Model
		FirstName string `gorm:"type:varchar;"`
		LastName  string `gorm:"type:varchar;"`
		FullName  string `gorm:"->;type:varchar GENERATED ALWAYS AS (first_name || ' ' || last_name) STORED"`
	}
	if err := DB.AutoMigrate(&exampleWithComputedCols{}); err != nil {
		t.Errorf("Failed test setup: error = %v", err)
	}

	original := exampleWithComputedCols{
		FirstName: "jon",
		LastName:  "hartman",
	}
	if err := DB.Create(&original).Error; err != nil {
		t.Errorf("Failed create: error = %v", err)
	}

	var result exampleWithComputedCols
	if err := DB.First(&result, original.ID).Error; err != nil {
		t.Errorf("Failed load: error = %v", err)
	}
	// It's expected that this is OK; since we're reloading a fresh struct
	if result.FullName != "jon hartman" {
		t.Errorf("Reloaded struct: computed 'Full Name' mismatch: have %s", result.FullName)
	}
	// However, this will fail - since the .Create() call above will not add full_name to the RETURNING clause (at least
	// with POSTGRES syntax). We get "RETURNING id, created_at, updated_at", while to make it return the computed column
	// value from 'first_name', it would need "RETURNING id, created_at, updated_at, full_name"
	if original.FullName != "jon hartman" {
		t.Errorf("Original struct on write: computed 'Full Name' mismatch: have %s", original.FullName)
	}
}
