package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type T1 struct {
	ID      uint
	Content string
}

type T2 struct {
	ID      uint
	Content string
}

type T3 struct {
	gorm.Model

	T1ID uint
	T1   *T1

	T2ID uint
	T2   *T2 `gorm:"-:migration"`
}

func TestGORM(t *testing.T) {
	_ = DB.Migrator().DropTable(&T1{}, &T2{}, &T3{})

	// Table t1 & t2 exists before
	err := DB.Exec(`CREATE TABLE t1 (
		id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
		content varchar(30) NOT NULL,
		PRIMARY KEY (id)
	);`).Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
		t.FailNow()
	}

	err = DB.Exec(`CREATE TABLE t2 (
		id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
		content varchar(30) NOT NULL,
		PRIMARY KEY (id)
	);`).Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
		t.FailNow()
	}

	// Migrate t3
	err = DB.AutoMigrate(&T3{})
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
		t.FailNow()
	}

	// t1 should not be modified
	columns, err := DB.Migrator().ColumnTypes(&T1{})
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
		t.FailNow()
	}

	for _, c := range columns {
		if c.Name() == "content" {
			columnType, _ := c.ColumnType()
			if columnType != "varchar(30)" {
				t.Errorf("T1: Expected type '%v', but got '%v'", "varchar(30)", columnType)
			}
		}
	}

	// t2 should not be modified
	columns, err = DB.Migrator().ColumnTypes(&T2{})
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
		t.FailNow()
	}

	for _, c := range columns {
		if c.Name() == "content" {
			columnType, _ := c.ColumnType()
			if columnType != "varchar(30)" {
				t.Errorf("T2: Expected type '%v', but got '%v'", "varchar(30)", columnType)
			}
		}
	}
}
