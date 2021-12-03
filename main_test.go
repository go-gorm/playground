package main

import (
	"testing"

	"gorm.io/datatypes"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	t.Run("json nil", func(t *testing.T) {
		user := User{
			Name: "jinzhu",
			Json: nil,
		}

		DB.Create(&user)

		var result User
		if err := DB.First(&result, user.ID).Error; err != nil {
			t.Errorf("Failed, got error: %v", err)
		}

		if result.Json != nil {
			t.Errorf("Failed, json was nil when saving but is '%v' afer reading'", result.Json)
		}
	})

	t.Run("json empty string", func(t *testing.T) {
		user := User{
			Name: "jinzhu",
			Json: datatypes.JSON{},
		}

		DB.Create(&user)

		var result User
		if err := DB.First(&result, user.ID).Error; err != nil {
			t.Errorf("Failed, got error: %v", err)
		}
	})

}
