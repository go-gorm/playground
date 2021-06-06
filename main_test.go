package main

import (
	"github.com/shopspring/decimal"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {

	DB.Create(&Product{Price: decimal.NewFromFloat(10.05)})
	DB.Create(&Product{Price: decimal.NewFromFloat(100.55)})

	var price decimal.Decimal

	DB.Model(&Product{}).Limit(1).Pluck("price", &price)
	if price.IntPart() == 0 {
		t.Error("Failed, got error")
	}

	var price2 decimal.Decimal

	DB.Model(&Product{}).Limit(1).Select("price").Scan(&price2)
	if price.IntPart() == 0 {
		t.Error("Failed, got error")
	}
}
