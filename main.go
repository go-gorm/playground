package main

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(postgres.Open("host=db user=user password=pass"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	type User struct {
		Name      string `gorm:"primaryKey"`
		UpdatedAt time.Time
	}

	user0 := User{Name: "Joe"}
	if err := db.Create(&user0).Error; err != nil {
		panic(err)
	}
	log.Printf("user0:  %v", user0)

	var user1 User
	if err := db.First(&user1).Error; err != nil {
		panic(err)
	}
	log.Printf("user1:  %v", user1)
}
