package main

import (

	"gorm.io/gorm"
)

// User has one `Account` (has one), many `Pets` (has many) and `Toys` (has many - polymorphic)
// He works in a Company (belongs to), he has a Manager (belongs to - single-table), and also managed a Team (has many - single-table)
// He speaks many languages (many to many) and has many friends (many to many - single-table)
// His pet also has one Toy (has one - polymorphic)

// Order .
type Order struct {
	gorm.Model
	CoinType             string `json:"coin_type" gorm:"type:varchar(32);Index;not null;default:'';comment:''"`
	Price                string `json:"price" gorm:"type:varchar(10);not null;default:'';comment:''"`
	Total                string `json:"total" gorm:"type:varchar(32);not null;default:0;comment:''"`
	Exchange             string `json:"exchange" gorm:"type:varchar(32);default:0;Index;comment:''"`
	ContractType         string `json:"contract_type" gorm:"type:varchar(32);Index;not null;default:'';comment:''"`
	AccountID            string `json:"account_id" gorm:"type:char(32);not null;default:'';comment:''"`
	TransactionAccountID string `json:"transaction_account_id" gorm:"type:char(32);not null;default:'';comment:''"`
	State                string `json:"state" gorm:"type:varchar(255);default:0;comment:''"`
}

type TransactionAccount struct {
	gorm.Model
	AccountID    string `json:"account_id" gorm:"type:char(32);Index;comment:''"`
	APIKey       string `json:"api_key" gorm:"type:varchar(255);comment:'';"`
	APISecret    string `json:"api_secret" gorm:"type:varchar(255);comment:''"`
	Exchange     string `json:"exchange" gorm:"type:varchar(32);default:0;Index;comment:''"`
	Passphrase   string `json:"passphrase"  gorm:"type:varchar(255);comment:''"`
	ContractType string `json:"contract_type" gorm:"type:varchar(32);Index;not null;default:'';comment:''"`
}

// StatisticsDetail .
type StatisticsDetail struct {
	TotalPrice float64 `json:"total_price"`
	CoinType   string  `json:"coin_type"`
	APIKey     string  `json:"api_key"`
	APISecret  string  `json:"api_secret"`
	Passphrase string  `json:"passphrase"`
	IDS        string  `json:"ids"`
}
