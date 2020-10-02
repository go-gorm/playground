package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	// I'll try to update a non existent user WITH PK, which shouldn't insert a new record
	// (Save update value in database, if the value doesn't have primary key, will insert it)
	user := User{Model: gorm.Model{ID: 1234567890}, Name: "somename"}
	DB.Save(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err == nil {
		t.Errorf("The user was created, which is not what the docs say.")
	}
}
