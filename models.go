package main

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

// User has one `Account` (has one), many `Pets` (has many) and `Toys` (has many - polymorphic)
// He works in a Company (belongs to), he has a Manager (belongs to - single-table), and also managed a Team (has many - single-table)
// He speaks many languages (many to many) and has many friends (many to many - single-table)
// His pet also has one Toy (has one - polymorphic)
type User struct {
	gorm.Model
	Name      string
	Age       uint
	Birthday  *time.Time
	Account   Account
	Pets      []*Pet
	Toys      []Toy `gorm:"polymorphic:Owner"`
	CompanyID *int
	Company   Company
	ManagerID *uint
	Manager   *User
	Team      []User     `gorm:"foreignkey:ManagerID"`
	Languages []Language `gorm:"many2many:UserSpeak"`
	Friends   []*User    `gorm:"many2many:user_friends"`
	Active    bool
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

type TestCase struct {
	gorm.Model `json:"-" yaml:"-"`
	ProductID  *uint  `gorm:"uniqueIndex:product_test,not null" json:"-" yaml:"-"`
	Product    string `gorm:"-" json:"product" yaml:"product"`
	Name       string `gorm:"type:varchar(255);uniqueIndex:product_test;not null" json:"testcase" yaml:"testcase"`
}

type Product struct {
	gorm.Model `json:"-" yaml:"-"`
	Name       string      `gorm:"type:varchar(100);uniqueIndex:product_name;not null" json:"name" yaml:"name"`
	TestCases  []*TestCase `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-" yaml:"-"`
}
