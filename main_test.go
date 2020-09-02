package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type Foo struct {
	ID    string `json:"id"`
	Value string `json:"value" gorm:"default:null"`
}

func TestGORM(t *testing.T) {
	foo := &Foo{
		ID:    "001",
		Value: "",
	}

	DB.Create(&foo)

	var result Foo
	if err := DB.First(&result, foo.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	
	if result.Value != "" {
		t.Errorf("Value is not empty")
	}
}
