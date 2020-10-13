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

	createErr := DB.Save(&UserWithJSON{}).Error
	if createErr != nil {
		t.Errorf("create Error : %v", createErr)
	}
}
