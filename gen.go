package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"playground/model"
)

var db *gorm.DB

func init() {
	db, _ = gorm.Open(mysql.Open("root:secret@(localhost:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
}

func generate() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./dal/query",
		Mode:    gen.WithDefaultQuery, /*WithQueryInterface, WithoutContext*/

		WithUnitTest: true,
	})
	g.UseDB(db)

	g.ApplyBasic(model.User{})

	g.Execute()
}
