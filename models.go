package main

import "gorm.io/gorm"

// User has one `Account` (has one), many `Pets` (has many) and `Toys` (has many - polymorphic)
// He works in a Company (belongs to), he has a Manager (belongs to - single-table), and also managed a Team (has many - single-table)
// He speaks many languages (many to many) and has many friends (many to many - single-table)
// His pet also has one Toy (has one - polymorphic)

const ossUrl = "http://www.baidu.com"

type User struct {
	Id       uint   `gorm:"primarykey"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Password string `json:"password"`
}

type UserInfo struct {
	Id       uint   `gorm:"primarykey"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

type Group struct {
	Id       uint   `gorm:"primarykey"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	Password string `json:"password"`
}

type GroupInfo struct {
	Id     uint   `gorm:"primarykey"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

func (v *UserInfo) AfterFind(db *gorm.DB) error {
	v.Avatar = ossUrl + v.Avatar
	return nil
}
func (v *Group) AfterFind(db *gorm.DB) error {
	return nil
}
func (v *GroupInfo) AfterFind(db *gorm.DB) error {
	v.Avatar = ossUrl + v.Avatar
	return nil
}
