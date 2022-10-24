package main

import (
	"gorm.io/gorm"
)

// User has one `Account` (has one), many `Pets` (has many) and `Toys` (has many - polymorphic)
// He works in a Company (belongs to), he has a Manager (belongs to - single-table), and also managed a Team (has many - single-table)
// He speaks many languages (many to many) and has many friends (many to many - single-table)
// His pet also has one Toy (has one - polymorphic)
type User struct {
	gorm.Model
	Name      string
	CompanyID *int
	Company   Company
	ManagerID *uint
	Manager   *User
	UserProps *UserProp `gorm:"foreignkey:CompanyID,ManagerID;references:CompanyID,ManagerID"`
}

type Company struct {
	ID   int
	Name string
}

type UserProp struct {
	ID        int
	CompanyID int
	ManagerID uint
	Value     string
}

// use this model in migrator to prevent fk creation with user's model
type UserPropMigration struct {
	ID        int
	CompanyID int
	Company   Company
	ManagerID uint
	Manager   User
	Value     string
}

func (r UserPropMigration) TableName() string {
	return "user_props"
}
