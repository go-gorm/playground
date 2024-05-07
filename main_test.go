package main

import (
	"errors"
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			t.Errorf("Failed, got error: %v", err)
		}
	}
}

type TableColumnWithoutDefault struct {
	gorm.Model

	Type int `gorm:"column:type;"`
}

func (TableColumnWithoutDefault) TableName() string {
	return "table_column_default"
}

type TableColumnWithDefault struct {
	gorm.Model

	Type int `gorm:"column:type;default:1;"`
}

func (TableColumnWithDefault) TableName() string {
	return "table_column_default"
}

func TestReMigrateColumnWithDefault(t *testing.T) {
	m1 := new(TableColumnWithoutDefault)
	if !DB.Migrator().HasTable(m1) {
		if err := DB.AutoMigrate(m1); err != nil {
			t.Fatalf("Failed to auto migrate, but got error %v\n", err)
		}
	}

	m2 := new(TableColumnWithDefault)
	if err := DB.AutoMigrate(m2); err != nil {
		t.Fatalf("Failed to auto migrate, but got error %v\n", err)
	}
}
