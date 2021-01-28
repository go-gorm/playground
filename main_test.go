package main

import (
	"gorm.io/gorm/clause"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	c1 := C{Name: "c1"}
	b1 := B{C: c1, Name: "b1"}
	a1 := A{B: b1, Name: "a1"}

	if err := DB.Create(&a1).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	// BUG: preloading all associations together with nested associations fails.
	tx := DB.
		Preload(clause.Associations).
		Preload("B.C")

	var res A
	if err := tx.Find(&res, a1.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
