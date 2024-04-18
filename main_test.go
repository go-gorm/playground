package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {

	type TestUser struct {
		gorm.Model
		Name      string `gorm:"index:name-idx"`
	}
	
	user := TestUser{
		Name: "jinzhu"
	}

	DB.Create(&user)

	indexes, err := DB.Migrator().GetIndexes(&user)
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
		return
	}
}
