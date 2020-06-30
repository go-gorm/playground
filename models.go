package main

import (
	"gorm.io/gorm"
	"log"
)

// User has one `Account` (has one), many `Pets` (has many) and `Toys` (has many - polymorphic)
// He works in a Company (belongs to), he has a Manager (belongs to - single-table), and also managed a Team (has many - single-table)
// He speaks many languages (many to many) and has many friends (many to many - single-table)
// His pet also has one Toy (has one - polymorphic)
type User struct {
	gorm.Model
	Avatar   string
	Nickname string
}

type Comment struct {
	gorm.Model
	UserId  uint
	User    *User `gorm:"foreignKey:UserId;AssociationForeignKey:Id"`
	Content string
}

func (m *User) AfterFind(db *gorm.DB) error {
	m.Avatar = "https://oss.xxx.com/" + m.Avatar
	log.Println("UserAfterFind")
	return nil
}

func (m *Comment) AfterFind(db *gorm.DB) error {
	log.Println("CommentAfterFind")
	return nil
}
