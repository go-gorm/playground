package main

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	// Insert
	DB.Transaction(func(tx *gorm.DB) error {
		process := ProcessTable{Name: "test"}
		if err := tx.Create(&process).Error; err != nil {
			t.Errorf("Failed process !!!!!!!, got error: %v", err)
			return err
		}
		return nil
	})

	// Ensure record does exist.
	k := &ProcessTable{}
	DB.First(k)
	t.Log("My Name is --------------------------", k.Name)
	var processes []ProcessTable

	// Try to Delete and return the record using Clauses(clause.Returning{})
	res := DB.Debug().Clauses(clause.Returning{}).Where(&ProcessTable{Name: "test"}).Delete(&processes)
	if res.Error != nil {
		t.Errorf("Failed process !!!!!!!, got error: %v", res.Error)
		return
	}
	if len(processes) == 0 {
		t.Errorf("Empty processes !!!!!!!")
		return
	}
	t.Log(processes[0].Name)
}
