package main

import (
	"gorm.io/gorm/clause"
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

	id := "s-1"
	val := "val1"
	if err := DB.Create(&SomeModel{
		SomeId: &id,
		SomeAtt: &SomeAttModel{
			Value: &val,
		},
	}).Error; err != nil {
		t.Errorf("failed to create SomeModel: %v", err)
	}

	result := &SomeModel{}
	if err := DB.Preload(clause.Associations).Take(result).Error; err != nil {
		t.Errorf("failed to get result: %v", err)
	}
	if result.SomeAtt == nil {
		t.Errorf("SomeAtt is nil")
	}
	if *result.SomeAtt.Value != val {
		t.Errorf("val is incorrect")
	}
}
