package main

import (
	"database/sql"

	"gorm.io/gorm"
)

type UserExt struct {
	EyeColour string
}

// User has one `Account` (has one), many `Pets` (has many) and `Toys` (has many - polymorphic)
// He works in a Company (belongs to), he has a Manager (belongs to - single-table), and also managed a Team (has many - single-table)
// He speaks many languages (many to many) and has many friends (many to many - single-table)
// His pet also has one Toy (has one - polymorphic)
type User[T any] struct {
	gorm.Model
	Name      string
	Languages []Language `gorm:"many2many:UserSpeak"`
	Extension UserExt    `gorm:"embedded"`
}

func (User[T]) TableName() string {
	return "users"
}

type Account struct {
	gorm.Model
	UserID sql.NullInt64
	Number string
}

type Pet struct {
	gorm.Model
	UserID *uint
	Name   string
	Toy    Toy `gorm:"polymorphic:Owner;"`
}

type Toy struct {
	gorm.Model
	Name      string
	OwnerID   string
	OwnerType string
}

type Company struct {
	ID   int
	Name string
}

type Language struct {
	Code string `gorm:"primarykey"`
	Name string
}
