package main

import (
	"gorm.io/gorm"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	// I have used Postgres for the test
	userWithAssignedID := User{
		Model: gorm.Model{ // first create with an assigned ID -> record is successfully inserted
			ID: 1,
		},
		Name: "jinzhu",
	}

	DB.Create(&userWithAssignedID)

	userWithoutAssignedID := User{
		Name: "jinzhu", // Since the ID 1 is already assigned to a record, ideally GORM should
	} // have checked and used the next available sequence value.
	err := DB.Create(&userWithoutAssignedID).Error
	if err != nil {
		t.Error(err.Error())
	}
}
