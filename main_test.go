package main

import (
	"testing"
	"time"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	DB.Use(&Plugin{})

	product := User{
		Name: "old joe",
		DOB:  time.Now(),
		CreditCards: []CreditCard{
			{
				Number: "987654321",
			},
		},
	}
	err := DB.Scopes(func(d *gorm.DB) *gorm.DB {
		return DB.Set("audited:current_user", "create-test")
	}).Save(&product).Error
	if err != nil {
		t.Errorf("unable to save: %s", err)
	}
	if product.CreatedBy == nil {
		t.Errorf("created_by for user not set")
	}
	if product.CreditCards[0].CreatedBy == nil {
		t.Errorf("created_by for credit card not set")
	}
}
