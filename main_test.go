package main

import (
	"testing"

	"github.com/shopspring/decimal"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	transactions := []Transaction{
		{Amount: decimal.NewFromFloat(17.99)},
		{Amount: decimal.NewFromFloat(17.99)},
		{Amount: decimal.NewFromFloat(17.99)},
		{Amount: decimal.NewFromFloat(17.99)},
		{Amount: decimal.NewFromFloat(17.99)},
	}

	for _, transaction := range transactions {
		DB.Create(&transaction)
	}

	var result decimal.Decimal

	err := DB.Table("transactions").Select(
		"SUM(amount)",
	).Row().Scan(&result)

	if err != nil {
		t.Errorf("'SELECT SUM(amount)' Failed, got error: %v", err)
	}

	if !result.Equals(decimal.NewFromFloat(89.95)) {
		t.Errorf("SUM(amount) is not 89.95, got %s", result)
	}
}
