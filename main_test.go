package main

import (
	"testing"
	"gorm.io/datatypes"
	"gorm.io/gorm"

)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: postgres

type UserWithJSON struct {
	gorm.Model
	Name       string
	Attributes datatypes.JSON
}

func TestGORM(t *testing.T) {
	err := DB.AutoMigrate(&UserWithJSON{}).Error

	if err != nil {
		t.Errorf("datatypes.json create error")
	}
}
