package main

import (
	"gorm.io/gorm"
)

// User has one `Account` (has one), many `Pets` (has many) and `Toys` (has many - polymorphic)
// He works in a Company (belongs to), he has a Manager (belongs to - single-table), and also managed a Team (has many - single-table)
// He speaks many languages (many to many) and has many friends (many to many - single-table)
// His pet also has one Toy (has one - polymorphic)
type User struct {
	gorm.Model

	ID           int64  `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Email        string `json:"email" gorm:"column:email"`
	Username     string `json:"username" gorm:"column:username;not null;unique"`
	PasswordHash string `json:"-" gorm:"column:password_hash;not null"`
	Version      string `json:"version" gorm:"column:version"`
}
