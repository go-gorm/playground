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
	tx := DB.Model(Catalog{})
	err := tx.Statement.Parse(tx.Statement.Model)
	assert.NoError(t, err)
	s := tx.Statement.Schema
	fmt.Println(s.Relationships.Relations)
	fmt.Println(s.Relationships.Relations["Scope"].FieldSchema.Relationships.Relations)
	fmt.Println(s.Relationships.Relations["Scope"].FieldSchema.Relationships.Relations["Labels"].FieldSchema.Relationships.Relations)
	fmt.Println(s.Relationships.Relations["Scope"].FieldSchema.Relationships.Relations["Labels"].FieldSchema.Relationships.Relations["CreatedBy"].FieldSchema.Relationships.Relations)

	assert.NotEmpty(t, s.Relationships.Relations["CreatedBy"], "Root CreatedBy missing relations")
	assert.NotEmpty(t, s.Relationships.Relations["Scope"].FieldSchema.Relationships.Relations["Labels"].FieldSchema.Relationships.Relations["CreatedBy"].FieldSchema.Relationships.Relations, "CreatedBy missing relations")
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

type Catalog struct {
	BaseModel
	Name  string
	Scope *Scope `gorm:"polymorphic:Owner;polymorphicValue:CampaignCatalog"`
}

type Scope struct {
	BaseModel
	OwnerID   int
	OwnerType string
	Labels    []*Label
}

type Label struct {
	BaseModel
	ScopeID int `gorm:"not null; index"`
	Scope   *Scope
}
