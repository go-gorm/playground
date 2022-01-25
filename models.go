package main

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Pets []*Pet
}

type Pet struct {
	gorm.Model
	Name     string
	FavToyID uint
	FavToy   *Toy
	UserID   uint
	User     *User
}

type Toy struct {
	gorm.Model
	Name  string
	PetID uint
	Pet   *Pet
}
