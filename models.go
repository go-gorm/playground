package main

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
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
	UUIDModel
	Name string
	DateTimeModel
}

type Language struct {
	Code string `gorm:"primarykey"`
	Name string
}

type CreationDateTimeModel struct {
	CreatedAt time.Time
}

func (m *CreationDateTimeModel) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Test creation")
	m.CreatedAt = time.Now()
	return nil
}

type DateTimeModel struct {
	CreationDateTimeModel
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (m *DateTimeModel) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Test datatime")
	m.UpdatedAt = time.Now()
	return nil
}

type UUIDModel struct {
	ID string
}

func isEmpty(id uuid.UUID) bool {
	for i := range id {
		if id[i] != 0 {
			return false
		}
	}
	return true
}

func (u *UUIDModel) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Test UUID")
	if len(u.ID) <= 0 {
		u.ID = uuid.New().String()
	}
	return nil
}
