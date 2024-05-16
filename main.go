package main

import (
	"fmt"
	"log"
)

type AB struct {
	Id  int `json:"id" gorm:"primaryKey"`
	BId int `json:"aId" gorm:"index"`
	AId int `json:"bId" gorm:"index"`
	B   B   `json:"b"`
	A   A   `json:"a"`
}

type A struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"uniqueIndex"`
}

type B struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"uniqueIndex"`
}

func main() {

	user := User{
		Account: Account{
			Number: "123456",
			Companies: []Company{
				{Name: "Corp1"}, {Name: "Corp2"},
			},
			Pet: Pet{
				Name: "Pet1",
			},
		},
	}
	DB.AutoMigrate(&User{}, &Account{}, &Pet{}, &Company{})
	DB.Create(&user)
	fmt.Println("-------------------------------------------------------")
	var count int64
	var result User
	DB = DB.Model(&User{}).
		Joins("Account").
		Joins("Account.Pet").
		Preload("Account.Companies")

	log.Println(count)
	log.Println(result)

}
