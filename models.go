package main

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
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

type Meeting struct {
	gorm.Model
	Id          string     `gorm:"type:uuid;primaryKey" json:"id"`
	IsLive      bool       `json:"isLive"`
	CompletedAt time.Time  `json:"completedAt"`
	Familiars   []Familiar `gorm:"many2many:familiar_meetings;"`
}

func (m *Meeting) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	m.Id = uuid.NewString()
	return
}

type Familiar struct {
	gorm.Model
	Id       string    `gorm:"type:uuid;primaryKey" json:"id"`
	Meetings []Meeting `gorm:"many2many:familiar_meetings;"`
	Name     string    `json:"name"`
}

func (f *Familiar) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	f.Id = uuid.NewString()
	return
}
