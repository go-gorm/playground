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
	Code string `gorm:"primaryKey"`
	Name string
}

type Department struct {
	TenantID       string     `gorm:"size:50;primaryKey"`
	DepartmentID   string     `gorm:"size:50;primaryKey"`
	DepartmentName string     `gorm:"size:255"`
	Employees      []Employee `gorm:"foreignKey:TenantID,DepartmentID;references:TenantID,DepartmentID"`
}

type Employee struct {
	TenantID     string `gorm:"size:50;primaryKey"`
	EmployeeID   string `gorm:"size:50;primaryKey"`
	EmployeeName string `gorm:"size:255"`
	DepartmentID string `gorm:"size:50"`
}

type Audit struct {
	AuditID   string `gorm:"size:50;primaryKey"`
	AuditDesc string `gorm:"size:255"`
}

// TableName setting
func (Employee) TableName() string {
	return "employee"
}

// TableName setting
func (Department) TableName() string {
	return "department"
}

// TableName setting
func (Audit) TableName() string {
	return "audit"
}
