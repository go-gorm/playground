package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{
		Name: "genofire",
		Manager: &User{
			Name: "jinzhu",
			Pets: []*Pet{
				&Pet{Name: "dog"},
				&Pet{Name: "cat"},
			},
		},
	}

	DB.Create(&user)

	var result User
	if err := DB.Preload("Manager.Pets").First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
