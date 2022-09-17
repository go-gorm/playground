package main

import (
	"gorm.io/playground/diff"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	if err := DB.AutoMigrate(&GameUser{}); err != nil {
		t.Fatal(err)
	}
	has := DB.Migrator().HasTable(&diff.GameUser{})
	if !has {
		t.Fatal("dont have table game_user")
	}
	tables, err := DB.Migrator().GetTables()
	if err != nil {
		t.Error(err)
	}
	t.Log("next migrate")

	if err := DB.AutoMigrate(&diff.GameUser{}); err != nil {
		t.Fatal(err)
	}
	t.Log("all tables in database")
	for _, table := range tables {
		t.Log(table)
	}
}
