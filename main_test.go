package main

import (
	"gorm.io/gorm"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: mysql

type Tag struct {
	gorm.Model
	OwnerID uint
	OwnerType string
}

type Foo struct {
	// An ID field solves the panic.
	// ID uint
	TagID uint
	Tag Tag `gorm:"polymorphic:Owner;polymorphicValue:foo"`
}

func TestGORM(t *testing.T) {
	if err := DB.AutoMigrate(&Tag{}, &Foo{}); err != nil {
		t.Error(err)
	}
}
