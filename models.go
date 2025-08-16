package main

import "github.com/shopspring/decimal"

type UserWallet struct {
	ID      uint64          `gorm:"column:id;primaryKey"`
	UserID  uint64          `gorm:"column:user_id;notNull"`
	Balance decimal.Decimal `gorm:"type:decimal(20,4)"`
}
