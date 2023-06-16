package main

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
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

type ValueDep struct {
	gorm.Model

	ID      int `gorm:"primarykey,autoIncrement"`
	ValueID int
	Name    string `gorm:"column:name; default:'default'"`
	Params  Params `gorm:"column:params; type:jsonb; default:'{}'"`
}

type Value struct {
	gorm.Model

	ID   int         `gorm:"primarykey,autoIncrement"`
	Deps []*ValueDep `gorm:"foreignKey:ValueID"`
}

type Params map[string]string

func (p *Params) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		return json.Unmarshal(v, p)
	case string:
		return json.Unmarshal([]byte(v), p)
	default:
		return fmt.Errorf("unsupported type: %T", v)
	}
}

func (p Params) Value() (driver.Value, error) {
	return json.Marshal(p)
}
