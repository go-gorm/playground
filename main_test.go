package main

import (
	"gorm.io/gorm"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	t.Cleanup(func() {
		err := DB.Exec("drop table tests").Error
		if err != nil {
			t.Fatalf("failed to drop table")
		}
	})

	columnTypes, err := DB.Migrator().ColumnTypes(&Test{})
	if err != nil {
		t.Fatalf("failed get columns")
	}

	columnCount1 := len(columnTypes)

	strCol := findColumn("some_str", columnTypes)
	if strCol == nil {
		t.Fatalf("some_str column not found")
	}

	nullable, _ := strCol.Nullable()
	if nullable {
		t.Fatalf("field is nullable")
	}

	err = DB.Migrator().AutoMigrate(&Test2{})
	if err != nil {
		t.Fatalf("automigrate failed")
	}

	columnTypes, err = DB.Migrator().ColumnTypes(&Test{})
	if err != nil {
		t.Fatalf("failed get columns")
	}

	columnCount2 := len(columnTypes)

	if columnCount2-columnCount1 != 1 {
		t.Fatalf("new column was not added")
	}

	strCol = findColumn("some_str", columnTypes)
	if strCol == nil {
		t.Fatalf("some_str column not found")
	}

	nullable, _ = strCol.Nullable()
	if !nullable {
		t.Fatalf("field is not nullable")
	}
}

func findColumn(name string, columns []gorm.ColumnType) gorm.ColumnType {
	for _, c := range columns {
		if c.Name() == name {
			return c
		}
	}

	return nil
}
