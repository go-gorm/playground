package main

import (
	"gorm.io/gorm"
)

type User1 struct {
	gorm.Model
	Email string
}

func (User1) TableName() string {
	return "users"
}

type User2 struct {
	gorm.Model
	Email string `gorm:"unique"`
}

func (User2) TableName() string {
	return "users"
}

type SomeNewTable struct {
	gorm.Model
	Email string `gorm:"unique"`
}

func (SomeNewTable) TableName() string {
	return "some_new_tables"
}
