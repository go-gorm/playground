package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type T1 struct {
	gorm.Model
	Content string `gorm:"type:varchar(30)"`
}

type T2 struct {
	gorm.Model
	T1ID uint
	T1   *T1
}

func TestGORM(t *testing.T) {
	// Table t1 exists before
	err := DB.Exec(`CREATE TABLE t1 (
		id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
		content text NOT NULL,
		PRIMARY KEY (id)
	);`).Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
		t.FailNow()
	}

	// Migrate t2
	err = DB.AutoMigrate(&T2{})
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
			if columnType != "text" {
				t.Errorf("Expected type '%v', but got '%v'", "text", columnType)
				t.FailNow()
			}
		}
	}
}
