package main

import (
	"testing"
	"github.com/google/uuid"
	"time"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: postgres

type UserWithUUID struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	TestID     *uuid.UUID `gorm:"type:uuid"`
}

func TestGORM(t *testing.T) {
	value := UserWithUUID{}
	err := DB.AutoMigrate(&value).Error

	if err != nil {
		t.Errorf("datatypes.json create error")
	}

	createErr := DB.Save(&value).Error
	if createErr != nil {
		t.Errorf("create Error : %v", createErr)
	}
}
