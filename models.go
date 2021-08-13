package main

import (
	"time"
)

type User struct {
	ID          string `gorm:"column:id;type:varchar(36);primary_key" json:"id"`
	Name        string
	LastLoginID *string  `gorm:"column:last_login_id;type:varchar(36);uniqueIndex"`
	LastLogin   *Login   `gorm:"foreignKey:LastLoginID"`
	Logins      []*Login `gorm:"foreignKey:UserID"`
}

type Login struct {
	ID       string `gorm:"column:id;type:varchar(36);primary_key" json:"id"`
	Location string
	UserID   string
	Time     *time.Time
}
