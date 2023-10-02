package main

import (
)

type Pet struct {
	FirstName string `gorm:"primaryKey"`
	LastName string `gorm:"primaryKey"`
	Tags []Tag
}

type Tag struct {
	PetFirstName string
	PetLastName string
	Name string
}
