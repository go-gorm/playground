package main

import (
	"database/sql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Pets      []Pet
	Languages []Language `gorm:"many2many:user_language;"`
}

type Language struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

type UserLanguage struct {
	UserID     uint `gorm:"primaryKey"`
	LanguageID uint `gorm:"primaryKey"`
	Language   Language
	Skilled    sql.NullBool
}

type Pet struct {
	gorm.Model
	UserID uint
	Name   string
}
