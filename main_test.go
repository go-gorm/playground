package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver


func TestGORM(t *testing.T) {
	corp := Corp{Base:Company{Name: "jinzhu"}}

	DB.Create(&corp)

	var result Corp
	if err := DB.First(&result, corp.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
