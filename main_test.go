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

// TestPostgresFullTextSearch tests the issue with @@ operator in Raw SQL
func TestPostgresFullTextSearch(t *testing.T) {
	// Skip for non-postgres drivers
	if DB.Dialector.Name() != "postgres" {
		t.Skip("Skipping test for non-postgres driver")
	}

	// Create test data
	user1 := User{Name: "search test user"}
	user2 := User{Name: "another user"}
	DB.Create(&user1)
	DB.Create(&user2)

	// Test 1: Using Raw with @@ operator and parameters - this will fail due to the issue
	t.Run("Raw with @@ operator and parameters", func(t *testing.T) {
		var users []User
		searchTerm := "search"

		// This query contains the @@ operator with parameters which causes issues
		err := DB.Raw(`
			SELECT * FROM users 
			WHERE ($1::text IS NULL OR to_tsvector('english', name) @@ plainto_tsquery('english', $1))
		`, searchTerm).Scan(&users).Error

		if err != nil {
			t.Errorf("Full-text search with @@ operator and parameters failed: %v", err)
		} else {
			t.Logf("Found %d users with full-text search", len(users))
			if len(users) == 0 {
				t.Error("Expected to find at least one user with full-text search")
			}
		}
	})

	// Test 2: Using Raw with @@ in SELECT clause - this will also fail
	t.Run("Raw with @@ in SELECT clause", func(t *testing.T) {
		var results []struct {
			ID      uint
			Name    string
			Matches bool
		}
		searchTerm := "search"

		// This query contains the @@ operator in the SELECT clause
		err := DB.Raw(`
			SELECT id, name, 
			       to_tsvector('english', name) @@ plainto_tsquery('english', $1) as matches
			FROM users
		`, searchTerm).Scan(&results).Error

		if err != nil {
			t.Errorf("Full-text search with @@ in SELECT clause failed: %v", err)
		} else {
			t.Logf("Found %d results with @@ in SELECT", len(results))
			for _, r := range results {
				t.Logf("User: %s, Matches: %v", r.Name, r.Matches)
			}
		}
	})

	// Test 3: Using Where with direct condition - this should work
	t.Run("Where with direct condition", func(t *testing.T) {
		var users []User

		err := DB.Where("name LIKE ?", "%search%").Find(&users).Error

		if err != nil {
			t.Errorf("Simple LIKE search failed: %v", err)
		} else {
			t.Logf("Found %d users with LIKE search", len(users))
			if len(users) == 0 {
				t.Error("Expected to find at least one user with LIKE search")
			}
		}
	})

	// Test 4: Using Where with full-text search - this should work
	t.Run("Where with full-text search", func(t *testing.T) {
		var users []User
		searchTerm := "search"

		// Using Where with the @@ operator should work
		err := DB.Where("to_tsvector('english', name) @@ plainto_tsquery('english', ?)", searchTerm).Find(&users).Error

		if err != nil {
			t.Errorf("Full-text search with Where failed: %v", err)
		} else {
			t.Logf("Found %d users with Where full-text search", len(users))
			if len(users) == 0 {
				t.Error("Expected to find at least one user with Where full-text search")
			}
		}
	})

	// Test 5: Workaround using ILIKE instead of @@ - this should work
	t.Run("ILIKE workaround", func(t *testing.T) {
		var users []User
		searchTerm := "search"

		// Using ILIKE as a workaround
		err := DB.Raw(`
			SELECT * FROM users 
			WHERE ($1::text IS NULL OR name ILIKE '%' || $1 || '%')
		`, searchTerm).Scan(&users).Error

		if err != nil {
			t.Errorf("ILIKE workaround failed: %v", err)
		} else {
			t.Logf("Found %d users with ILIKE workaround", len(users))
			if len(users) == 0 {
				t.Error("Expected to find at least one user with ILIKE workaround")
			}
		}
	})

	// Cleanup
	DB.Unscoped().Delete(&user1)
	DB.Unscoped().Delete(&user2)
}
