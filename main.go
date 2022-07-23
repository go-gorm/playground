package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Table struct {
	Id   uint
	Test bool `gorm:"default:1"`
}

func main() {
	dsn := "gorm:gorm@tcp(mysql)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&Table{})

	createTestRecord := Table{Test: false}
	db.Create(&createTestRecord)
}
