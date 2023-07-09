package main

import (
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	UserId uint
}

func (Address) TableName() string {
	return "addresses"
}

// Represents the edited address that has a uniqueness constraint
// so that no user has more than one address
type Address2 struct {
	gorm.Model
	UserId      uint `gorm:"uniqueIndex:,where: deleted_at is null"`
	SecondField string
}

func (Address2) TableName() string {
	return "addresses"
}

// User has one address.
// Represent this same user with two structs to mock running a database migration
// multiple times.
type User struct {
	gorm.Model
	Name    string
	Address Address `gorm:"foreignKey:user_id"`
}
