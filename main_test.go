package main

import (
	"testing"

	"gorm.io/gorm/clause"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite

func TestGORM(t *testing.T) {
	users := []map[string]any{
		{"Name": "FOOBAR"},
	}
	// Without batching all works fine
	// err := DB.Model(&User{}).Clauses(clause.OnConflict{DoNothing: true}).Create(&users).Error
	// if err != nil {
	// 	t.Errorf("Failed, got error: %v", err)
	// }
	// With bathing it doesn't
	err := DB.Model(&User{}).Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(&users, 100).Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	var result User
	if err := DB.First(&result).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
