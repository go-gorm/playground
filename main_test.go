package main

import (
	"testing"
	"gorm.io/gorm/clause"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	db := DB.Clauses(clause.Locking{Strength: "UPDATE"})
	var parents []*Parent
	err := db.Preload("Children").Find(&parents).Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	err = db.Delete(&Child{}, "parent_id = 5").Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
