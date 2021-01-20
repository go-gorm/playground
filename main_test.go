package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Model: gorm.Model{
		ID: 1,
	},
		Name: "zzc9402",
	}

	DB.Save(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if result.Age != 1 {
		t.Errorf("failed,age should be 1, not %d", result.Age)
	}

}
