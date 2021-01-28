package main

import (
	"gorm.io/gorm"
)


type A struct {
	gorm.Model
	BB []B `gorm:"many2many:a_bs"`
	Name string
}

type B struct {
	gorm.Model
	CID uint
	C C
	Name string
}

type C struct {
	gorm.Model
	Name string
}