package main

import (
	"gorm.io/gorm/clause"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	thing := Thing{
		SomeID:  "1234",
		OtherID: "1234",
		Data:    "something",
	}

	DB.Create(&thing)

	thing2 := Thing{
		SomeID:  "1234",
		OtherID: "1234",
		Data:    "something else",
	}

	result := DB.Clauses(clause.OnConflict{
		OnConstraint: "something_idx",
		UpdateAll:    true,
	}).Create(&thing2)
	if result.Error != nil {
		t.Errorf("creating second thing: %v", result.Error)
	}

	var things []Thing
	if err := DB.Find(&things).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if len(things) > 1 {
		t.Errorf("expected 1 thing got more")
	}
}
