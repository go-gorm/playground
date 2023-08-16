package main

import (
	"gorm.io/gen"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func createTable() error {
	return DB.Exec("CREATE TABLE IF NOT EXISTS example(example TEXT NOT NULL DEFAULT '^\\S+$')").Error
}

func TestGORM(t *testing.T) {
	if err := createTable(); err != nil {
		t.Fatalf("failed to create table, got error: %v", err)
	}
	g := gen.NewGenerator(gen.Config{
		Mode: gen.WithDefaultQuery,
	})
	g.UseDB(DB)
	g.ApplyBasic(g.GenerateModel("example"))
	g.Execute()
}
