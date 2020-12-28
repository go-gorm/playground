package main

import (
	"gorm.io/gorm"
)

// Example from https://gorm.io/docs/belongs_to.html#Override-References
type User struct {
	gorm.Model
	Name      string
	CompanyID string
	Company   Company `gorm:"references:Code"` // use Code as references
}

type Company struct {
	ID   int
	Code string
	Name string
}