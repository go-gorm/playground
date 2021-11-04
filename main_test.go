package main

import (
	"testing"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{
		Name: "jinzhu",
		Company: Company{
			Name: "acme",
		},
		Manager: &User{
			Name: "joe",
		},
	}

	DB.Create(&user)

	unscopedPreload := func(db *gorm.DB) *gorm.DB {
		return db.Unscoped()
	}

	// soft delete company and manager
	if result := DB.Where("name = ?", "acme").Delete(&Company{}); result.Error != nil {
		t.Errorf("Error when deleting Company: %v", result.Error)
	}

	if result := DB.Where("name = ?", "joe").Delete(&User{}); result.Error != nil {
		t.Errorf("Error when deleting User: %v", result.Error)
	}

	// This test succeeded as expected
	t.Run("Distinct preloads", func(t *testing.T) {
		var u User
		if result := DB.Preload("Company", unscopedPreload).
			Preload("Manager", unscopedPreload).
			Find(&u, user.ID); result.Error != nil {
			t.Errorf("Find failed: %v", result.Error)
		}
		if u.Company.Name != "acme" {
			t.Errorf("Wrong company name")
		}
		if u.Manager == nil {
			t.Errorf("Wrong manager")
		} else if u.Manager.Name != "joe" {
			t.Errorf("Wrong manager name")
		}
	})

	// This test fails, which is unexpected
	t.Run("Preload all", func(t *testing.T) {
		var u User
		if result := DB.Preload(clause.Associations, unscopedPreload).
			Find(&u, user.ID); result.Error != nil {
			t.Errorf("Find failed: %v", result.Error)
		}
		if u.Company.Name != "acme" {
			t.Errorf("Wrong company name")
		}
		if u.Manager == nil {
			t.Errorf("Wrong manager")
		} else if u.Manager.Name != "joe" {
			t.Errorf("Wrong manager name")
		}
	})
}
