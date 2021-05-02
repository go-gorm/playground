package main

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	DeletedBy string
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

func (u *User) BeforeDelete(tx *gorm.DB) error {
	// how to update u.DeletedBy of all records being deleted
	// by this transaction on BeforeDelete?

	// I tried to modify the current structure:
	u.DeletedBy = "gabriel"

	// I tried to use SetColumn()
	tx.Statement.SetColumn("DeletedBy", "gabriel")
	tx.Statement.SetColumn("deleted_by", "gabriel")

	// I tried to use AddClause()
	tx.Statement.AddClause(clause.Set{{Column: clause.Column{Name: "DeletedBy"}, Value: "gabriel"}})
	tx.Statement.AddClause(clause.Set{{Column: clause.Column{Name: "deleted_by"}, Value: "gabriel"}})

	// if I use tx.Update() here, I don't know how to match the
	// same records being deleted by this transaction
	// tx.Model(&User{}).Where("???").Update("DeletedBy", "gabriel")

	return nil
}
