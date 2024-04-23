package main

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    string
	Company *Company
}

type Company struct {
	ID        int
	Name      string
	UserID    uint
	Addresses []Addresses
}

type Addresses struct {
	ID        int
	Name      string
	CompanyID uint
	Company   Company
}
