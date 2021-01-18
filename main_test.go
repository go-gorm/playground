package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	tx := DB.Model(TestUser{})
	err := tx.Statement.Parse(tx.Statement.Model)
	assert.NoError(t, err)
	s := tx.Statement.Schema
	fmt.Println(s.Relationships.Relations)
	fs := s.Relationships.Relations["CreatedBy"].FieldSchema
	fmt.Println(fs.Relationships.Relations)

	fmt.Println(s.Name, " -> ", fs.Name)

	if _, ok := fs.Relationships.Relations["CreatedBy"]; !ok {
		t.Errorf("Missing CreatedBy")
	}
}

type TestUser struct {
	BaseModel

	Name string
}

type BaseModel struct {
	gorm.Model
	CreatedByID *int
	CreatedBy   *TestUser
}
