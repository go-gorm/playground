package main

import (
	"testing"
	"time"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver


type Test struct {
	Model
}

type Model struct {
	ID        uint `gorm:"type:serial"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}


func TestGORM(t *testing.T) {
	DB.AutoMigrate(&Test{})
	
	if err := DB.Create(&Test{}).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	if err := DB.Create(&Test{}).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

}
