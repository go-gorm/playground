package main

import (
	"gorm.io/gorm"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	country := Country{Name: "Test"}
	address1 := Address{
		Country: country,
	}
	org := Org{
		Address1: address1,
		Address2: address1,
	}

	DB.Create(&org)

	stmt := &gorm.Statement{DB: DB}
	stmt.Parse(&org)

	belongsTo := stmt.Schema.Relationships.BelongsTo
	embedded := stmt.Schema.Relationships.EmbeddedRelations
	relations := stmt.Schema.Relationships.Relations

	if len(belongsTo) != 2 {
		t.Errorf("Expected 2 belongsTo, got %v", len(belongsTo))
	}
	if len(embedded) != 2 {
		t.Errorf("Expected 2 embedded, got %v", len(embedded))
	}
	if len(relations) != 2 {
		t.Errorf("Expected 2 relations, got %v", len(relations))
	}
}
