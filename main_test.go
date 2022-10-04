package main

import (
	"fmt"
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: v1.23.10
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func migrateT1() error {
	type MyTable struct {
		gorm.Model
		FieldOne string
	}

	return DB.AutoMigrate(&MyTable{})
}

func migrateT2() error {

	type MyTable struct {
		NewFieldTwo string
	}

	return DB.AutoMigrate(&MyTable{})
}

func migrateT3() error {
	type MyTable struct {
		NewFieldThree string
	}
	return DB.AutoMigrate(&MyTable{})
}

func TestGormMigrationWithPreparedStatements(t *testing.T) {
	fmt.Println("Migrating T1")
	err := migrateT1()
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	fmt.Println("Migrating T2")

	err = migrateT2()
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	fmt.Println("Migrating T3")
	err = migrateT3()
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

}
