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
	DB.AutoMigrate(&value)

	createErr := DB.Save(&value).Error
	if createErr != nil {
		t.Errorf("create Error : %v", createErr)
	}
}
