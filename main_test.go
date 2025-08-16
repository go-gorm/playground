package main

import (
	"fmt"
	"testing"

	"gorm.io/gorm/clause"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {

	if err := DB.AutoMigrate(&Item{}); err != nil {
		fmt.Println("Failed to auto-migrate:", err)
		return
	}

	items := []Item{
		{Name: "Item1"},
		{Name: "Item2"},
		{Name: "Item3"},
		{Name: "Item4"},
		{Name: "Item5"},
		{Name: "Item6"},
	}

	if err := DB.Model(&Item{}).
		Omit(clause.Associations).
		Clauses(clause.OnConflict{
			DoUpdates: clause.AssignmentColumns([]string{"updated_at"}),
		}).
		Create(&items).Error; err != nil {
		fmt.Println("Item creation failed:", err)
		return
	}

	for _, item := range items {
		fmt.Println("Initial insert ID:", item.ID)
	}

	newItems := []Item{
		{Name: "Item1"},
		{Name: "Item2"},
		{Name: "Item7"},
		{Name: "Item3"},
		{Name: "Item4"},
		{Name: "Item8"},
	}

	if err := DB.Model(&Item{}).
		Omit(clause.Associations).
		Clauses(clause.OnConflict{
			DoUpdates: clause.AssignmentColumns([]string{"updated_at"}),
		}).
		Create(&newItems).Error; err != nil {
		fmt.Println("Item creation failed:", err)
		return
	}

	for _, item := range newItems {
		fmt.Println("After upsert ID:", item.ID)
	}

	var actualItems []Item
	if err := DB.Find(&actualItems).Error; err != nil {
		fmt.Println("Failed to fetch items:", err)
		return
	}

	fmt.Println("Real IDs from DB:")
	for _, item := range actualItems {
		fmt.Println("ID:", item.ID, ", Name:", item.Name)
	}

	expectedIDs := []uint64{1, 2, 3, 4, 5, 6, 7, 8}
	for i, actualItem := range actualItems {
		if actualItem.ID != expectedIDs[i] {
			t.Errorf("Unexpected ID for %s: got %d, want %d", actualItem.Name, actualItem.ID, expectedIDs[i])
		}
	}
}
