package main

import (
	"fmt"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	orders := []Order{
		{
			CoinType:             "ADA",
			Price:                "2",
			Total:                "5",
			AccountID:            "1",
			TransactionAccountID: "1",
			Exchange:             "a",
			ContractType:         "aa",
			State:                "-1",
		},
		{
			CoinType:             "ADA",
			Price:                "4",
			Total:                "10",
			AccountID:            "1",
			TransactionAccountID: "1",
			Exchange:             "a",
			ContractType:         "aa",
			State:                "-1",
		},
		{
			CoinType:             "bbc",
			Price:                "3",
			Total:                "6",
			AccountID:            "2",
			TransactionAccountID: "2",
			Exchange:             "a",
			ContractType:         "aa",
			State:                "-1",
		},
	}
	
	DB.Create(&orders)
	
	account := []TransactionAccount{
		{
			AccountID:    "1",
			APIKey:       "11",
			APISecret:    "11",
			Exchange:     "a",
			Passphrase:   "11",
			ContractType: "aa",
		},
		{
			AccountID:    "2",
			APIKey:       "22",
			APISecret:    "22",
			Exchange:     "a",
			Passphrase:   "22",
			ContractType: "aa",
		},
	}
	DB.Create(&account)
	var (
		exchange     = "a"
		contractType = "aa"
		state        = "-1"
	)
	
	var result StatisticsDetail
	if err := DB.Model(&Order{}).Debug().
		Select("SUM( total ) as total_price,`orders`.coin_type,(transaction_accounts.api_key) as api_key,"+
			"(transaction_accounts.api_secret) as api_secret,(transaction_accounts.passphrase) as passphrase,"+
			"GROUP_CONCAT(`orders`.transaction_account_id) as ids").
		Joins("left JOIN transaction_accounts on transaction_accounts.id = `orders`.transaction_account_id").
		Where("`orders`.exchange =? AND `orders`.contract_type=? AND `orders`.state=?", exchange, contractType, state).
		Group("`orders`.transaction_account_id,`orders`.coin_type").Scan(&result).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	fmt.Printf("%#v\n", result)
}
