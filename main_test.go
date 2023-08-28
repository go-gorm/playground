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

	if err := DB.AutoMigrate(Toy{}); err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if err := DB.AutoMigrate(Toy{}); err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
