package main

import (
	"testing"

	"gorm.io/gorm/clause"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	DB.Create(&User{Name: "jinzhu", Active: true})
	DB.Create(&User{Name: "ezequiel", Active: true})
	DB.Create(&User{Name: "jinzhu", Active: false})
	DB.Create(&User{Name: "ezequiel", Active: false})

	var results []User
	conds := clause.Not(clause.And(
		clause.Eq{
			Column: clause.Column{Name: "name"},
			Value:  "jinzhu",
		},
		clause.Eq{
			Column: clause.Column{Name: "active"},
			Value:  true,
		},
	))
	// Query with `NOT(name = 'jinzhu' AND active = true)` should return last three records
	// By De Morgan's Law this is equivalent to `name != 'jinzhu' OR active != true`
	// but current version (since 1.25.6) transforms this to
	// `name != 'jinzhu' AND active != true` which returns only the last record

	err := DB.
		Model(&User{}).
		Clauses(clause.Where{Exprs: []clause.Expression{conds}}).
		Find(&results).
		Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	if len(results) != 3 {
		t.Errorf("Failed, expected 3 records, got %v", len(results))
	}
	for _, r := range results {
		if r.Name == "jinzhu" && r.Active {
			t.Errorf("Failed, unexpected record: %v", r)
		}
	}

	conds = clause.Not(clause.Or(
		clause.Eq{
			Column: clause.Column{Name: "name"},
			Value:  "jinzhu",
		},
		clause.Eq{
			Column: clause.Column{Name: "active"},
			Value:  true,
		},
	))
	// Query with `NOT(name = 'jinzhu' OR active = true)` should return last record
	err = DB.
		Model(&User{}).
		Clauses(clause.Where{Exprs: []clause.Expression{conds}}).
		Find(&results).
		Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	if len(results) != 1 {
		t.Errorf("Failed, expected 1 records, got %v", len(results))
	}
	for _, r := range results {
		if r.Name == "jinzhu" || r.Active {
			t.Errorf("Failed, unexpected record: %v", r)
		}
	}
}
