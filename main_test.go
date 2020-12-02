package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{
		UnitID: 8,
		ID:     1,
		Name:   "Andrei",
		Projects: []Project{
			{
				ID:    1,
				Title: "First project",
			},
			{
				ID:    2,
				Title: "Second project",
			},
		},
	}

	if err := DB.Create(&user).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	var result User
	if err := DB.Where(&User{UnitID: user.UnitID, ID: user.ID}).First(&result).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	var projects []Project
	if err := DB.Where(&Project{UnitID: user.UnitID, UserID: user.ID}).Find(&projects).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	if len(projects) != len(user.Projects) {
		t.Errorf("Failed, projects length should be %d, got %d", len(user.Projects), len(result.Projects))
	}
}
