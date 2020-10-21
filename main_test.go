package main

import (
	"os"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// MYSQL_REPO: https://github.com/asmeikal/mysql.git
// POSTGRES_REPO: https://github.com/asmeikal/postgres.git
// GORM_BRANCH: master
// TEST_DRIVERS: mysql, postgres, sqlserver

func TestColumnTypeLengthInfo(t *testing.T) {
	columns, err := DB.Migrator().ColumnTypes(&User{})
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	for _, c := range columns {
		if c.Name() == "string_len" {
			if _, ok := c.Length(); !ok {
				t.Errorf("%s: missing length information on column %s", os.Getenv("GORM_DIALECT"), c.Name())
			}
		}
	}
}

func TestColumnTypeNullableInfo(t *testing.T) {
	columns, err := DB.Migrator().ColumnTypes(&User{})
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	for _, c := range columns {
		if _, ok := c.Nullable(); !ok {
			t.Errorf("%s: missing nullable information on column %s", os.Getenv("GORM_DIALECT"), c.Name())
		}
	}
}
