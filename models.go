package main

import (
	"time"

	"gorm.io/gorm"
)

// User has many projects
type User struct {
	UnitID    uint `gorm:"primaryKey"`
	ID        uint `gorm:"primaryKey;autoIncrement:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name     string
	Projects []Project `gorm:"foreignKey:UnitID,UserID"`
}

type Project struct {
	UnitID    uint `gorm:"primaryKey"`
	ID        uint `gorm:"primaryKey;autoIncrement:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	UserID uint
	Title  string
}
