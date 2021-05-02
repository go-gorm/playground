package main

import (
	"database/sql"
	"time"
	"fmt"
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
	CreatedBy string `gorm:"not null;default:null"`
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
	gorm.Model
	ID   int
	Name string
}

type Language struct {
	gorm.Model
	Code string
	Name string
}

type UserLanguage struct {
	gorm.Model
	UserID int
	LanguageID int
	CreatedBy string `gorm:"not null;default:null"`
}

func (ul *User) BeforeCreate(tx *gorm.DB) error {
	if user, ok := tx.Get("user"); ok {
		ul.CreatedBy = fmt.Sprintf("%v", user)
	}
	return nil
}

func (ul *UserLanguage) BeforeCreate(tx *gorm.DB) error {
	// why I can't get "user" here?
	if user, ok := tx.Get("user"); ok {
		ul.CreatedBy = fmt.Sprintf("%v", user)
	}
	return nil
}
