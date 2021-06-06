package main

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Price decimal.Decimal `gorm:"type:decimal(20,4)"`
}
