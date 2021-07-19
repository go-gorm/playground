package main

import (
	"database/sql"
	"os"
	"path"
	"path/filepath"
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
	Birthday  time.Time
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

// My Models

type Model struct {
	Id        int64     `gorm:"type:bigserial;primaryKey" json:"id"`
	CreatedAt time.Time `sql:"default:NOW()" json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Organisation struct {
	Model
	Name      string        `json:"name"`
	WmbusKeys *WmbusKeys    `gorm:"type:jsonb" json:"wmbusKeys"`
	LogoId    sql.NullInt64 `json:"logoFileId"`
	Logo      File          `gorm:"foreignKey:LogoId" json:"logo"`
}

type WmbusKey struct {
	Id  string `json:"id"`
	Key string `json:"key"`
}

type WmbusKeys struct {
	Keys []WmbusKey `json:"keys"`
}

var DataPath = "data/"

type File struct {
	Model
	Name     string `json:"name"` // Some sane name to identify the file
	Path     string `json:"path"` // Path to file on filesystem relative to data folder
	MimeType string `json:"mimeType"`
}

// PathAbs returns the absolute path on filesystem
func (f *File) PathAbs() string {
	if f == nil {
		return ""
	}
	wd, _ := os.Getwd()
	return filepath.ToSlash(path.Join(wd, DataPath, f.Path))
}

// PathUrl returns the URL that can be used to fetch the File from the http server
func (f *File) PathUrl() string {
	if f == nil {
		return ""
	}
	return filepath.ToSlash(path.Join(DataPath, f.Path))
}
