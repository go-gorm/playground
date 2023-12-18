package main

import (
	"gorm.io/gorm"
)

type A struct {
	gorm.Model
	Name string `gorm:"unique"`
}

type B struct {
	gorm.Model
	Name string `gorm:"unique"`
}

type Good struct {
	A   A
	AId uint `gorm:"primaryKey;autoIncrement:false"`
	B   B
	BId uint `gorm:"primaryKey;autoIncrement:false"`
}

type Bad struct {
	A   A
	AId uint `gorm:"primaryKey,priority:2;autoIncrement:false"`
	B   B
	BId uint `gorm:"primaryKey,priority:1;autoIncrement:false"`
}
