package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	company := Company{
		Name: "jinzhu",
		MyEnum: Enum_ONE,
	}

	DB.Create(&company)

	var result Company
	if err := DB.First(&result, company.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
