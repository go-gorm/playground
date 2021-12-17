package main

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FavPetID uint
	FavPet   *Pet
}

type Pet struct {
	gorm.Model
	UserID *uint
	User   *User
}
