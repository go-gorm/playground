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

type ProcessTable struct {
	ID         uint64    `gorm:"column:f_id;primaryKey;autoIncrement;not null;comment:'主键'"`
	Index      string    `gorm:"column:f_index;type:char(36);uniqueIndex;not null;comment:'UUID4'"`
	Name       string    `gorm:"column:f_name;type:varchar(128);uniqueIndex;not null;comment:''"`
	CreateTime time.Time `gorm:"column:f_create_time;autoCreateTime;not null;comment:''"`
	UpdateTime time.Time `gorm:"column:f_update_time;autoUpdateTime;not null;comment:''"`
}
