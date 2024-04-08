package main

type Common struct {
	TenantID uint `gorm:"primarykey"`
}

type User struct {
	Common
	UserID string `gorm:"primarykey"`
	Name   string
}
