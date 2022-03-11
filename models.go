package main

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Amount decimal.Decimal `json:"amount" gorm:"type:DECIMAL(20,8);"`
}
