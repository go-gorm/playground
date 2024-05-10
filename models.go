package main

import (
	"database/sql"

	"gorm.io/gorm"
)

// User has one `Account` (has one), many `Pets` (has many) and `Toys` (has many - polymorphic)
// He works in a Company (belongs to), he has a Manager (belongs to - single-table), and also managed a Team (has many - single-table)
// He speaks many languages (many to many) and has many friends (many to many - single-table)
// His pet also has one Toy (has one - polymorphic)
type User struct {
	gorm.Model
	Name    string
	Account Account
}

type Account struct {
	gorm.Model
	UserID    sql.NullInt64
	Number    string
	Companies []Company
	Pet       Pet
}

type Pet struct {
	gorm.Model
	UserID    *uint
	AccountID *uint
	Name      string
	Toy       Toy `gorm:"polymorphic:Owner;"`
}

type Toy struct {
	gorm.Model
	Name      string
	OwnerID   string
	OwnerType string
}

type Company struct {
	ID        int
	AccountID int32
	Name      string
}

type Language struct {
	Code string `gorm:"primarykey"`
	Name string
}
