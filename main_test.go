package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	categories := []Category{{LabelID: 1, CategoryID: 1}, {LabelID: 1, CategoryID: 2}, {LabelID: 1, CategoryID: 3}, {LabelID: 2, CategoryID: 1}, {LabelID: 2, CategoryID: 2}, {LabelID: 2, CategoryID: 3}, {LabelID: 3, CategoryID: 1}, {LabelID: 3, CategoryID: 2}, {LabelID: 3, CategoryID: 3}}

	DB.Create(&categories)

	var counter int
	var categorys []Category
	counter = 0
	result := DB.Unscoped().Debug().
		FindInBatches(&categorys, 5, func(tx *gorm.DB, batch int) error {
			for _, category := range categorys {
				if category.CategoryID > 0 {
					counter = counter + 1
				}
			}
			return nil
		})
	if result.Error != nil {
		t.Errorf("FindInBatches Error = %s", result.Error)
		t.Fail()
	}
	if counter != len(categories) {
		t.Errorf("number of records returned incorrect, expected %d received %d", len(categories), counter)
		t.Fail()
	}
}
