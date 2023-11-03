package main

import (
	"github.com/shopspring/decimal"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	wallet := UserWallet{UserID: 1, Balance: decimal.NewFromFloat(99.60)}

	if err := DB.Create(&wallet).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	var balance decimal.Decimal
	if err := DB.Model(&UserWallet{}).Limit(1).Pluck("balance", &balance).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if err := DB.Model(&UserWallet{}).Limit(1).Pluck("balance", balance).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
