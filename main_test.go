package main

import (
	"testing"

	"gorm.io/gorm/clause"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{
		Name: "jinzhu",
		UUID: "testuuid",
		Aux: &UserAux{
			Aux: "TestVal",
		},
	}

	if err := DB.Preload(clause.Associations).Create(&user).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	if user.UUID != user.Aux.UUID {
		t.Errorf("aux doesnot match")
	}

}
