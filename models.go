package main

import (
	"database/sql"
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

type IDType uint
type Code string

type BrokenUpdate struct {
	ID        IDType         `gorm:"primaryKey;<-:create" json:"id"`
	CreatedAt time.Time      `gorm:"<-:create" json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	JobName         string    `gorm:"uniqueIndex;->;type:varchar GENERATED ALWAYS AS ('thing-session-' || CAST(id as TEXT)) STORED;" json:"jobName"`
	SecretToken     string    `gorm:"unique" json:"-"`
	Code            Code      `gorm:"uniqueIndex" json:"code"`
	LastCheckinTime time.Time `json:"-"`
	EmbeddedStruct  `gorm:"embeddedPrefix:embed_"`
	StringMember    string `json:"stringMember"`
}

type EmbeddedStruct struct {
	Foo  string `json:"foo"`
	Bar  string `json:"bar"`
	Quux string `json:"quux"`
}

func (s *BrokenUpdate) BeforeCreate(db *gorm.DB) (err error) {
	s.SecretToken = "generated"
	s.Code = "generated"
	s.LastCheckinTime = time.Now()
	return nil
}

func (s *BrokenUpdate) AfterCreate(db *gorm.DB) (err error) {
	s.JobName = fmt.Sprintf("thing-session-%d", s.ID)

	return nil
}

func (s *BrokenUpdate) BeforeDelete(tx *gorm.DB) (err error) {
	s.EmbeddedStruct = EmbeddedStruct{}
	tx = tx.Model(s).Select("*").Updates(*s)
	return tx.Error
}
