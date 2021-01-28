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
	c2 := C{Name: "c2"}

	b1 := B{C: c1, Name: "b1"}
	b2 := B{C: c2, Name: "b2"}
	b3 := B{C: c1, Name: "b3"}

	a1 := A{BB: []B{b1,b2}, Name: "a1"}
	a2 := A{BB: []B{b3}, Name: "a2"}

	aa := []A{a1,a2}

	if err := DB.Create(&aa).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	// BUG: preloading all associations together with nested associations fails.
	tx := DB.Where(A{Name: a1.Name}).
		Preload(clause.Associations).
		Preload("BB.C")

	var res A
	if err := tx.Find(&res).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
