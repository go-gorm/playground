package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"reflect"
	"time"

	"gorm.io/gorm"
)

// User has one `Account` (has one), many `Pets` (has many) and `Toys` (has many - polymorphic)
// He works in a Company (belongs to), he has a Manager (belongs to - single-table), and also managed a Team (has many - single-table)
// He speaks many languages (many to many) and has many friends (many to many - single-table)
// His pet also has one Toy (has one - polymorphic)
type User struct {
	gorm.Model
	Name        string
	Age         uint
	Birthday    *time.Time
	Account     Account
	Pets        []*Pet
	Toys        []Toy `gorm:"polymorphic:Owner"`
	CompanyID   *int
	Company     Company
	ManagerID   *uint
	Manager     *User
	Team        []User     `gorm:"foreignkey:ManagerID"`
	Languages   []Language `gorm:"many2many:UserSpeak"`
	Friends     []*User    `gorm:"many2many:user_friends"`
	Active      bool
	ExternalKey interface{} `gorm:"type:varchar (36)""`
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

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(30-10+1) + 10
	user.ExternalKey = fmt.Sprintf("%d", random)
	return
}

func (user *User) AfterFind(tx *gorm.DB) (err error) {
	//user.ExternalKey is **interface{}
	tempID := reflect.Indirect(reflect.Indirect(reflect.ValueOf(user.ExternalKey))).Interface()
	external, ok := tempID.([]uint8)
	if ok {
		user.ExternalKey = string(external)
		return
	}
	user.ExternalKey = tempID

	return
}
