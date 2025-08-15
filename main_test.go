package main

import (
	"fmt"
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

func TestBooleanDefaultBug(t *testing.T) {
	// Auto migrate the Tree model
	if err := DB.AutoMigrate(&Tree{}); err != nil {
		t.Fatalf("Failed to migrate Tree model: %v", err)
	}

	// Clean up any existing data
	DB.Where("1 = 1").Delete(&Tree{})

	// Create a tree with IsAlive explicitly set to false
	tree := Tree{
		IsAlive: false,
		Height:  10.0,
	}

	if err := DB.Create(&tree).Error; err != nil {
		t.Fatalf("Failed to create tree: %v", err)
	}

	// Query the database to check the actual stored values
	var debugCount struct {
		IsAliveTrue  int64 `json:"is_alive_true"`
		IsAliveFalse int64 `json:"is_alive_false"`
	}

	DB.Model(&Tree{}).Where("is_alive = ?", true).Count(&debugCount.IsAliveTrue)
	DB.Model(&Tree{}).Where("is_alive = ?", false).Count(&debugCount.IsAliveFalse)

	fmt.Printf("DEBUG: After insert - has %d trees with is_alive=true, %d with is_alive=false\n",
		debugCount.IsAliveTrue, debugCount.IsAliveFalse)

	// The bug: even though we set IsAlive to false, it might be stored as true due to the default value
	if debugCount.IsAliveFalse == 0 {
		t.Errorf("BUG REPRODUCED: Expected 1 tree with is_alive=false, but found %d. The boolean field with default:true cannot be set to false", debugCount.IsAliveFalse)
	}

	// Verify by reading the record back
	var retrievedTree Tree
	if err := DB.First(&retrievedTree, tree.ID).Error; err != nil {
		t.Fatalf("Failed to retrieve tree: %v", err)
	}

	t.Logf("Original tree.IsAlive: %v", tree.IsAlive)
	t.Logf("Retrieved tree.IsAlive: %v", retrievedTree.IsAlive)

	if retrievedTree.IsAlive != false {
		t.Errorf("BUG REPRODUCED: Expected tree.IsAlive to be false, but got %v", retrievedTree.IsAlive)
	}
}
