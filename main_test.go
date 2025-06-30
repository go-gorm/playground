package main

import (
	"slices"
	"testing"
)

func TestPreloadWithOrCondition(t *testing.T) {
	// --- Setup ---
	DB.Migrator().DropTable(&Dataset{}, &DatasetColumn{})
	if err := DB.AutoMigrate(&Dataset{}, &DatasetColumn{}); err != nil {
		t.Fatalf("Failed to migrate database: %v", err)
	}

	// Create sample data for the target dataset
	dataset := Dataset{
		Name: "My Dataset",
		Columns: []*DatasetColumn{
			{Name: "System Column 1", IsSystem: true, IsHidden: false},       // Should be loaded
			{Name: "Visible Column 2", IsSystem: false, IsHidden: false},     // Should be loaded
			{Name: "Hidden System Column 3", IsSystem: true, IsHidden: true}, // Should be loaded
		},
	}
	if err := DB.Create(&dataset).Error; err != nil {
		t.Fatalf("Failed to create sample data: %v", err)
	}

	dataset2 := Dataset{
		Name: "My Dataset 2",
	}
	if err := DB.Create(&dataset2).Error; err != nil {
		t.Fatalf("Failed to create sample data: %v", err)
	}

	// Create a column that SHOULD NOT be loaded but will be by the buggy query
	// It belongs to a different dataset but is visible.
	// The correct query would filter by dataset_id first.
	// The buggy query will include it because `is_hidden = false` is true.
	rogueColumn := DatasetColumn{
		Name:      "Rogue Visible Column",
		DatasetID: dataset2.ID, // Different dataset
		IsSystem:  false,
		IsHidden:  false,
	}
	if err := DB.Create(&rogueColumn).Error; err != nil {
		t.Fatalf("Failed to create rogue column: %v", err)
	}

	// --- Reproduce the Bug ---
	var loadedDataset Dataset
	err := DB.Model(&Dataset{}).
		Where(&Dataset{ID: dataset.ID}).
		Scopes(Preload("Columns", QueryColumnsFilter())).
		First(&loadedDataset).Error

	if err != nil {
		t.Fatalf("Error during preload: %v", err)
	}

	// --- Assertion ---
	var loadedNames []string
	for _, col := range loadedDataset.Columns {
		loadedNames = append(loadedNames, col.Name)
	}

	// This is the key assertion that will now fail.
	// The buggy query will incorrectly load "Rogue Visible Column"
	// because the OR condition bypasses the dataset_id check.
	if slices.Contains(loadedNames, "Rogue Visible Column") {
		t.Errorf("FATAL BUG: Loaded a column from another dataset ('Rogue Visible Column') due to incorrect WHERE clause logic")
	}

	expectedCount := 3
	if len(loadedDataset.Columns) != expectedCount {
		// This assertion might also fail now, as the count will be 4 instead of 3.
		t.Errorf("Expected to load %d columns for dataset %d, but got %d. Loaded names: %v",
			expectedCount, dataset.ID, len(loadedNames), loadedNames)
	}
}
