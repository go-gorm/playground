package main

import (
	"testing"
)

type Embed struct {
	Status string `gorm:"not null;default:''"`
}

type Bugs struct {
	Embed

	ID  int64
	Foo string
}

func TestGORM(t *testing.T) {
	if err := DB.AutoMigrate(new(Bugs)); err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if err := DB.Select("foo").Create(new(Bugs)).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
