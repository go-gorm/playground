package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/plugin/soft_delete"
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


type BasicGorm struct {
	ID        int64 `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt soft_delete.DeletedAt `gorm:"index"`
}


type PortalDb struct {
	BasicGorm
	Name        string
	Description string
	Creator     string
}

func (PortalDb) TableName() string {
	return "pg_portals"
}

func (r *PortalDb) GetDb(ctx context.Context) *gorm.DB {
	return DB.Unscoped().Where("deleted_at is null or deleted_at = 0").WithContext(ctx).Model(r)
}

func (r *PortalDb) Get(ctx context.Context) error {
	if e := r.GetDb(ctx).Model(r).Where("id = ?", r.ID).First(r).Error; e != nil {
		fmt.Println(r.TableName()+" Get error: %+v", e)
		return errors.New("查询门户配置失败")
	}
	return nil
}
