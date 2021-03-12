package main

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

// PersonOne * one model without an ignored property
type PersonOne struct {
	ID             int `gorm:"primaryKey"`
	SourcePersonID *int
	Friends        []PersonOne `gorm:"foreignKey:SourcePersonID;references:ID"`
	//AnotherPro    interface{} `gorm:"-"`
}

// PersonTwo * one model with an ignored property
type PersonTwo struct {
	ID             int `gorm:"primaryKey"`
	SourcePersonID *int
	Friends        []PersonTwo `gorm:"foreignKey:SourcePersonID;references:ID"`
	AnotherPro     interface{} `gorm:"-"`
}

func PersonOneEqual(p1, p2 PersonOne) bool {
	if p1.ID != p2.ID {
		return false
	}

	if p1.SourcePersonID != nil && p2.SourcePersonID != nil {
		if *p1.SourcePersonID != *p2.SourcePersonID {
			return false
		}
	} else if p1.SourcePersonID != nil || p2.SourcePersonID != nil {
		return false
	}

	if len(p1.Friends) != len(p2.Friends) {
		return false
	}
	for i := 0; i < len(p1.Friends); i++ { // TODO
		if !PersonOneEqual(p1.Friends[i], p2.Friends[i]) {
			return false
		}
	}

	return true
}

func PersonTwoEqual(p1, p2 PersonTwo) bool {
	if p1.ID != p2.ID {
		return false
	}

	if p1.SourcePersonID != nil && p2.SourcePersonID != nil {
		if *p1.SourcePersonID != *p2.SourcePersonID {
			return false
		}
	} else if p1.SourcePersonID != nil || p2.SourcePersonID != nil {
		return false
	}

	if len(p1.Friends) != len(p2.Friends) {
		return false
	}
	for i := 0; i < len(p1.Friends); i++ { // TODO
		if !PersonTwoEqual(p1.Friends[i], p2.Friends[i]) {
			return false
		}
	}

	return true
}

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
