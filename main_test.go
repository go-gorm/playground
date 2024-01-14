package main

import (
	"fmt"
	"sync"
	"testing"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
	"gorm.io/gorm/utils/tests"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	var users []User
	var id uint = 1
	// Works with a value
	if err := DB.Table("users").Where("id = ?", id).Find(&users).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	// Works with a pointer
	if err := DB.Table("users").Where("id = ?", &id).Find(&users).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	ids := []uint{1, 2, 3}
	// Works with a slice
	if err := DB.Table("users").Where("id IN (?)", ids).Find(&users).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	// Does not work with a slice pointer
	if err := DB.Table("users").Where("id IN (?)", &ids).Find(&users).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}

// From gorm/clause/expression_test.go
func TestExpr(t *testing.T) {
	ids := []uint{1, 2, 3}
	results := []struct {
		SQL    string
		Result string
		Vars   []interface{}
	}{{
		SQL:    "SELECT * FROM users WHERE users.id IN (?)",
		Vars:   []interface{}{ids},
		Result: "SELECT * FROM users WHERE users.id IN (?,?,?)", // Correctly expanded
	}, {
		SQL:    "SELECT * FROM users WHERE users.id IN (?)",
		Vars:   []interface{}{&ids},
		Result: "SELECT * FROM users WHERE users.id IN (?,?,?)", // Expected this to expand aswell
	}}

	for idx, result := range results {
		t.Run(fmt.Sprintf("case #%v", idx), func(t *testing.T) {
			user, _ := schema.Parse(&tests.User{}, &sync.Map{}, DB.NamingStrategy)
			stmt := &gorm.Statement{DB: DB, Table: user.Table, Schema: user, Clauses: map[string]clause.Clause{}}
			clause.Expr{SQL: result.SQL, Vars: result.Vars}.Build(stmt)
			if stmt.SQL.String() != result.Result {
				t.Errorf("generated SQL is not equal, expects %v, but got %v", result.Result, stmt.SQL.String())
			}
		})
	}
}
