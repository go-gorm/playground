package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: postgres

type Event struct {
	gorm.Model
	ID  string `gorm:"primaryKey"`
	UID uint32 `gorm:"not null;autoIncrement"`
}

func TestGORM(t *testing.T) {
	if err := DB.AutoMigrate(&Event{}); err != nil {
		t.Errorf("Failed can't migrate pgx to gorm [1]")
	}
	if err := DB.AutoMigrate(&Event{}); err != nil {
		t.Errorf("Failed can't migrate pgx to gorm [2]")
	}
}
