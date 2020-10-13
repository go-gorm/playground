package main

import (
	"testing"
	"github.com/satori/go.uuid"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: postgres

type UserWithJSON struct {
	TestID     *uuid.UUID `gorm:"index;type:uuid"`
}

func TestGORM(t *testing.T) {
	err := DB.AutoMigrate(&UserWithJSON{}).Error

	if err != nil {
		t.Errorf("datatypes.json create error")
	}

	err = DB.Create(&UserWithJSON{}).Error
	if err != nil {
		t.Errorf("create Error : %v", err)
	}
}
