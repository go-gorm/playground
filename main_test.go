package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type OtherSchemaTable struct {
	AColumn string `gorm:"primary_key"`
}

func (OtherSchemaTable) TableName() string {
	return "otherschema.other_schema_table"
}

func TestGORM(t *testing.T) {
	if err := DB.AutoMigrate(&OtherSchemaTable{}); err != nil {
		t.Errorf("Failed to automigrate: %s", err)
	}

	result := make([]OtherSchemaTable, 0)

	if err := DB.Find(&result).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
