package main

import (
	"github.com/google/uuid"
	"time"

	"gorm.io/gorm"
)

// User has one `Account` (has one), many `Pets` (has many) and `Toys` (has many - polymorphic)
// He works in a Company (belongs to), he has a Manager (belongs to - single-table), and also managed a Team (has many - single-table)
// He speaks many languages (many to many) and has many friends (many to many - single-table)
// His pet also has one Toy (has one - polymorphic)
type User struct {
	ID        string `gorm:"type:varchar(100);primaryKey"` // A unique identifier for the rule
	Name      string
	Age       uint
	Birthday  *time.Time
	Account   Account
	Pets      []*Pet
	Toys      []Toy `gorm:"polymorphic:Owner"`
	CompanyID *string `gorm:"foreignKey:CompanyID"`
	Company   Company
	ManagerID *string
	Manager   *User
	Team      []User     `gorm:"foreignkey:ManagerID"`
	Languages []Language `gorm:"many2many:UserSpeak"`
	Friends   []*User    `gorm:"many2many:user_friends"`
	Active    bool
}

func (q *User) BeforeCreate(tx *gorm.DB) (err error) {
	q.ID = uuid.New().String()
	return
}

type Account struct {
	gorm.Model
	UserID *string
	Number string
}

type Pet struct {
	gorm.Model
	UserID *string
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
	ID   string `gorm:"type:varchar(100);primaryKey"` // A unique identifier for the rule
	Name string
}

func (q *Company) BeforeCreate(tx *gorm.DB) (err error) {
	q.ID = uuid.New().String()
	return
}

type Language struct {
	Code string `gorm:"primarykey"`
	Name string
}
