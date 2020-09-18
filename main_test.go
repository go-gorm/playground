package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type TestTable struct {
	Name     string `gorm:"column:name;uniqueIndex:test_tables_name_nick_name_uind"`
	NickName string `gorm:"column:nick_name;uniqueIndex:test_tables_name_nick_name_uind"`
}

func TestGORM(t *testing.T) {
	DB.Exec(`
	CREATE TABLE public.test_tables (
		"name" text NULL
	)`)

	if err := DB.AutoMigrate(&TestTable{}); err != nil {
		t.Fail()
	}
	if !DB.Migrator().HasIndex(&TestTable{}, "test_tables_name_nick_name_uind") {
		t.Fail()
	}
}
