package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type MultiTable struct {
	Field1 string `gorm:"primarykey"`
	Field2 string `gorm:"primarykey"`
	Field3 string `gorm:"primarykey"`
}

func TestGORM(t *testing.T) {

	if err := DB.AutoMigrate(&MultiTable{}); err != nil {
		t.Fatalf("couldnt create multi_table")
	}

	DB.Exec("delete from multi_tables")

	rec := MultiTable{
		Field1: "val1",
		Field2: "val2",
		Field3: "val3",
	}

	if err := DB.Create(&rec).Error; err != nil {
		t.Error(err)
	}

	if err := DB.Delete(&rec).Error; err != nil {
		t.Error(err)
	}

}
