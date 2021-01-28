package main

import (
	"gorm.io/gorm"
)

type A struct {
	gorm.Model
	B B
	Name string
}

type B struct {
	gorm.Model
	AID uint
	C C
	Name string
}

type C struct {
	gorm.Model
	BID uint
	Name string
}