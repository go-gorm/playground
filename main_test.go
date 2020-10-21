package main

import (
	"fmt"
	"gorm.io/gorm"
	"os"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// MYSQL_REPO: https://github.com/asmeikal/mysql.git
// POSTGRES_REPO: https://github.com/asmeikal/postgres.git
// GORM_BRANCH: master
// TEST_DRIVERS: mysql, postgres, sqlserver

func getColumn(value interface{}, name string) (gorm.ColumnType, error) {
	columns, err := DB.Migrator().ColumnTypes(&User{})
	if err != nil {
		return nil, fmt.Errorf("failed to get columns for entity %v: %w", value, err)
	}

	for _, c := range columns {
		if c.Name() == name {
			return c, nil
		}
	}

	return nil, fmt.Errorf("column %s not found", name)
}

func TestHasTable(t *testing.T) {
	out := DB.Migrator().HasTable(&User{})

	if !out {
		t.Errorf("table for %#v not found", &User{})
	}
}

func TestHasColumn(t *testing.T) {
	out := DB.Migrator().HasColumn(&User{}, "name")

	if !out {
		t.Errorf("column name not found")
	}
}

func TestColumnTypeLengthInfo(t *testing.T) {
	columnsWithLength := []string{
		"string_len",
	}

	for _, columnName := range columnsWithLength {
		t.Run(fmt.Sprintf("column %s has length", columnName), func(t *testing.T) {
			column, err := getColumn(&User{}, columnName)

			if err != nil {
				t.Errorf("Failed, got error: %v", err)
			}

			if _, ok := column.Length(); !ok {
				t.Errorf("%s: missing length information on column %s", os.Getenv("GORM_DIALECT"), column.Name())
			}
		})
	}
}

func TestColumnTypeNullableInfo(t *testing.T) {
	columnsWithNullable := []string{
		"name",
		"age",
		"string_len",
	}

	for _, columnName := range columnsWithNullable {
		t.Run(fmt.Sprintf("column %s has nullable", columnName), func(t *testing.T) {
			column, err := getColumn(&User{}, columnName)

			if err != nil {
				t.Errorf("Failed, got error: %v", err)
			}

			if _, ok := column.Nullable(); !ok {
				t.Errorf("%s: missing nullable information on column %s", os.Getenv("GORM_DIALECT"), column.Name())
			}
		})
	}
}
