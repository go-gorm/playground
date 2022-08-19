package main

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code   string
	Price  uint
	Author Author `gorm:"type:JSON;serializer:json"`
}

type Author struct {
	Name  string
	Email string
}
