package main

import (
	"fmt"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{FirstName: "jinzhu", LastName: "Dop", Email: "radnom@gmail.com", Id: 0}
	fmt.Println("This is User ", user)

	err1 := DB.Model(&User{}).Create(&user)
	if err1 != nil {
		fmt.Println("Error while saving", err1.Error)
		t.Errorf("Failed, got error no 1: %v", err1)
	}

	var result User
	if err := DB.First(&result, user.Id).Error; err != nil {
		fmt.Println("Error while getting", err.Error())
		t.Errorf("Failed, got error: %v", err)
	}
}
