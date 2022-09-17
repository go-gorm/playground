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

	t.Log("next migrate")
	t.Log("just a simulator. actually two migrate use the same one struct")

	if err := DB.AutoMigrate(&diff.GameUser{}); err != nil {
		t.Fatal(err)
	}
}
