package main

import (
	"gorm.io/gorm/clause"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	users := []User{
		{Name: "jinzhu", Age: 18},
		{Name: "jinzhu1", Age: 18},
	}

	DB.Model(&users).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "name"}},
			DoUpdates: clause.AssignmentColumns([]string{"age", "name"}),
		}, clause.Returning{}).CreateInBatches(&users, len(users))

	var result User
	if err := DB.First(&result, users[0].ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
