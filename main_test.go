package main

import (
	"sync"
	"testing"

	"gorm.io/gorm/schema"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}

func TestParseJSONBFieldWithDefaultValue(t *testing.T) {
	type Payload struct {
		ID string
	}
	type Table struct {
		ID      string  `gorm:"primaryKey;type:uuid"`
		Payload Payload `gorm:"serializer:json;type:jsonb;default:'{}'"`
		Str     string  `gorm:"default:'default'"`
	}

	s, err := schema.Parse(&Table{}, &sync.Map{}, schema.NamingStrategy{})
	if err != nil {
		t.Fatalf("Unable to parse schema: %v", err)
	}
	if s.FieldsByDBName["payload"].DefaultValue != "{}" {
		t.Errorf("Expected default value for payload to be `{}`, got %q", s.FieldsByDBName["payload"].DefaultValue)
	}
	if s.FieldsByDBName["str"].DefaultValue != "default" {
		t.Errorf("Expected default value for str to be `default`, got %q", s.FieldsByDBName["str"].DefaultValue)
	}
}
