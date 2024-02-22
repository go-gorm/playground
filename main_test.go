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
		Model: gorm.Model{
			ID: 1,
		},
		Name: "jinzhu",
		Pets: []*Pet{
			{
				Model: gorm.Model{
					ID: 1,
				},
				UserID: &[]uint{1}[0],
				Toy: Toy{
					OwnerID: 1,
				},
			},
		},
	}

	tcs := []struct {
		name             string
		shouldHardDelete bool
	}{
		{
			name:             "soft delete",
			shouldHardDelete: false,
		},
		{
			name:             "hard delete",
			shouldHardDelete: true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			tx := DB.Begin()
			tx.Create(&user)

			if tc.shouldHardDelete {
				tx = tx.Unscoped()
			}

			if err := tx.Select(clause.Associations).Delete(&user).Error; err != nil {
				t.Errorf("Failed, got error: %v", err)
			}

			du := User{Model: gorm.Model{ID: 1}}
			if err := tx.Find(&du).Error; err != nil {
				t.Errorf("Failed, got error: %v", err)
			}

			tx.Rollback()
		})
	}

}
