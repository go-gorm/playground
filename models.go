package main

type Entity struct {
	ID               int
	Name             string `gorm:"uniqueIndex;size:64"`
}
