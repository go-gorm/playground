package main

import (
	"fmt"

	"gorm.io/gorm"
)

func main() {
	user := User{Name: "jinzhu"}
	DB.Create(&user)
	var ids []int
	result := DB.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Table("users")
	}).Pluck("id", &ids)
	if result.Error != nil {
		fmt.Printf("err: %v", result.Error)
	} else {
		fmt.Printf("ids: %v", ids)
	}
}
