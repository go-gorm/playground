package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"column:username;not null;" json:"username"`
	Status   int    `gorm:"column:status;not null;default:1" json:"status"`
}

func (User) TableName() string {
	return "gorm_test_user_tmp"
}
