package main

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Profiles []Profile `gorm:"many2many:user_profiles;foreignKey:Refer;joinForeignKey:UserReferID;References:UserRefer;JoinReferences:UserRefer"`
	Refer    uint
}

type Profile struct {
	gorm.Model
	Name      string
	UserRefer uint
}