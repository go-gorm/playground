package main

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	//Active bool
}
