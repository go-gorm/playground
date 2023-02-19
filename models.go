package main

import (
	"gorm.io/gorm"
)

type UserFollow struct {
	gorm.Model
	UserID      int64
	FollowedUID int64
}

func (UserFollow) TableName() string {
	return "user_follow"
}
