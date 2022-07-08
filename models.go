package main

import (
	// "database/sql"
	// "time"

	// "gorm.io/gorm"
)

// User has one `Account` (has one), many `Pets` (has many) and `Toys` (has many - polymorphic)
// He works in a Company (belongs to), he has a Manager (belongs to - single-table), and also managed a Team (has many - single-table)
// He speaks many languages (many to many) and has many friends (many to many - single-table)
// His pet also has one Toy (has one - polymorphic)
type User struct {
	Name   string `gorm:"primaryKey"`
	Gender int    `gorm:"primaryKey"`
	Cond   string
	// fields2 string
}
