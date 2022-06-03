package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	tx := DB.Begin()
	tx2 := tx.Begin()
	tx2.Create(&user)

	var result User
	if err := tx2.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	if err := tx2.Commit().Error; err != nil {
		t.Errorf("Failed to commit: %v", err)
	}
	tx.Commit()
}
