package main

import (
	"testing"

	"gorm.io/gorm/clause"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	products := []Product{{Name: "test-product"}, {Name: "test-product"}}

	for _, product := range products {
		res := DB.Omit(clause.Associations).
			Clauses(clause.OnConflict{
				Columns: []clause.Column{
					{Name: "name"},
				},
				DoUpdates: clause.AssignmentColumns([]string{
					"name",
				}),
			}).
			Create(&product)

		if err := res.Error; err != nil {
			t.Error(err)
		}

		if product.ID == 0 {
			t.Errorf("product: %s has an ID: %d", product.Name, product.ID)
		}
	}
}
