package main

import (
	"log"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: postgres

func TestGORM(t *testing.T) {
	user := User{FirstName: "jinzhu"}

	DB.Create(&user)

	result := User{}
	if err := DB.FirstOrCreate(&result, User{
		Email: "test@mail.com",
	}).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	log.Println(result)
}
