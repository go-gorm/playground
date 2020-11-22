package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	type foo struct {
		gorm.Model
		BarCID string
	}

	// Migrate the schema
	err := DB.AutoMigrate(&foo{})
	if err != nil {
		t.Fatal("failed to migrate database")
	}

	// Create
	err = DB.Create(&foo{BarCID: "some value"}).Error
	if err != nil {
		t.Fatal("failed to create model")
	}

	var m foo

	// Read with expected column name
	err = DB.First(&m, "bar_cid = ?", "some value").Error // find with BarCID "some value"
	if err != nil {
		t.Fatal("can't fetch with expected column name")
	}

	// Read with actual column name
	err = DB.First(&m, "bar_c_id = ?", "some value").Error // find with BarCID "some value"
	if err == nil {
		t.Fatal("can fetch with unexpected column name")
	}
}
