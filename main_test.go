package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}

type Sample struct {
	ID     string `gorm:"size:128;default:(uuid())" json:"id"`
	Status string `gorm:"size:256; default:''; not null" json:"status"`
}

func TestStringID(t *testing.T) {

	err := DB.AutoMigrate(&Sample{})
	if err != nil {
		t.Errorf("unable to auto-migrate: %v", err)
	}

	w := Sample{Status: "12345"}

	err = DB.Create(&w).Error
	if err != nil {
		t.Errorf("unable to insert: %v", err)
	}

	if w.ID == "" {
		t.Errorf("w:ID should be uuid")
	}

}
