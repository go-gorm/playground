package main

import (
	"gorm.io/gorm"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	userID := uint(1024)
	user := User{
		Model: gorm.Model{
			ID: userID,
		},
		Name: "jinzhu",
		Pets: []*Pet{
			{UserID: &userID, Name: "foo"},
			{UserID: &userID, Name: "bar"},
		},
	}

	DB.Session(&gorm.Session{FullSaveAssociations: true}).Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
