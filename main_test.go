package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	language := Language{Code: "pl-PL", Name: "Polish"}

	DB.Create(&language)

	if err := DB.Delete(Language{}, language.Code).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
