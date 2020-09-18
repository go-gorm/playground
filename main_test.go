package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type TestData struct {
	Id   int
	Data string
}

func TestGORM(t *testing.T) {
	tx := DB.Begin()
	tx.Migrator().AutoMigrate(&TestData{})
	if err = tx.Rollback().Error; err != nil {
		t.Error(err)
	}

	test := TestData{Data: "blub"}
	if err = DB.Save(&test).Error; err == nil {
		t.Log("table shouldn't exist")
		t.FailNow()
	}	
}
