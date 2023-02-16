package main

import (
	"testing"

	"gorm.io/gorm/clause"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestClauseChaining(t *testing.T) {
	DB.Create(&Company{
		ID: 1,
		Name: "foo",
	})

	DB.Create(&Company{
		ID: 2,
		Name: "bar",
	})

	var results []Company
	DB.Clauses(&clause.Expr{
		SQL: `(id = 1 OR id = 2)`,
	}, &clause.Expr{
		SQL: `name = 'bar'`,
	}).Find(&results)

	if l := len(results); l != 1 {
		t.Fatalf("expected 1, got %d", l)
	}
}
