package main

import (
	"testing"

	"gorm.io/gorm"
	"gorm.io/playground/models"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := models.User{Name: "jinzhu"}

	DB.Create(&user)

	var result models.User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}

// TestPreloadCallbackCloneValue tests that the *gorm.DB passed to Preload callback
// allows safe reuse of the db variable with chain methods.
//
// This pattern worked in v1.26.x but breaks in v1.30.0+ because the db passed
// to the callback has clone=0, causing chain methods to modify the same instance.
func TestPreloadCallbackCloneValue(t *testing.T) {
	user := models.User{
		Name: "preload-test-user",
		Pets: []*models.Pet{
			{Name: "cat1"},
			{Name: "dog1"},
			{Name: "cat2"},
		},
	}
	if err := DB.Create(&user).Error; err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	var result models.User
	err := DB.Preload("Pets", func(db *gorm.DB) *gorm.DB {
		// Build two separate conditions using the same db variable.
		// If db has clone > 0, each Where() returns a new instance (safe).
		// If db has clone = 0, each Where() modifies the same instance (unsafe).
		catQuery := db.Where("name LIKE ?", "cat%")
		dogQuery := db.Where("name LIKE ?", "dog%")

		return db.Where(catQuery.Or(dogQuery))
	}).First(&result, user.ID).Error

	if err != nil {
		t.Fatalf("failed to query: %v", err)
	}

	// Expected: All 3 pets (cat1, dog1, cat2) should be loaded
	if len(result.Pets) != 3 {
		t.Errorf("expected 3 pets, got %d", len(result.Pets))
		for _, pet := range result.Pets {
			t.Logf("  loaded pet: %s", pet.Name)
		}
	}
}
