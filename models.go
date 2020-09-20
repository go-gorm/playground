package main

type User struct {
	UserID uint    `gorm:"primarykey"`
	Groups []Group `gorm:"many2many:user_groups;foreignKey:UserID;joinForeignKey:GroupID"`
	Name   string
}

type Group struct {
	GroupID uint `gorm:"primarykey"`
	Name    string
}
