package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	// user := User{Name: "jinzhu"}

	// DB.Create(&user)

	// var result User
	// if err := DB.First(&result, user.ID).Error; err != nil {
	// 	t.Errorf("Failed, got error: %v", err)
	// }
	//var users []User
	var user User
	multiUser := make([]User, 2)
	toy := make([]Toy, 2)
	toy1 := make([]Toy, 3)
	multiUser[0].Name = "hukejie"
	multiUser[1].Name = "brother"
	toy[0].Name = "toms"
	toy[1].Name = "jack"

	toy1[0].Name = "ronad"
	toy1[1].Name = "mess"
	multiUser[0].Toys = toy
	multiUser[1].Toys = toy1
	result := DB.Create(&multiUser)
	if result.Error != nil {
		t.Errorf("Failed, createMultiUser error: %v", result.Error)
	}

	// DB.Save(&multiUser[0])
	// time.Sleep(time.Second * 10)
	// DB.Save(&multiUser[1])

	//time.Sleep(time.Second * 60)

	multiUser[0].Name = "hukejie of test"
	multiUser[1].Name = "sister"

	toy[0].Name = "tony"
	toy[1].Name = "jackd"

	toy1[0].Name = "c7"
	toy1[1].Name = "ber"

	multiUser[0].Toys = toy
	multiUser[1].Toys = toy1
	result = DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&multiUser[0])
	if result.Error != nil {
		t.Errorf("Failed, updatesMultiUser[0] error: %v", result.Error)
	}

	// time.Sleep(time.Second * 30)
	result = DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&multiUser[1])
	if result.Error != nil {
		t.Errorf("Failed, updatesMultiUser[1] error: %v", result.Error)
	}

	// time.Sleep(time.Second * 30)

	multiUser[0].Name = "test1"
	multiUser[1].Name = "test sister"

	toy[0].Name = "tonytest"
	toy[1].Name = "jackdtest"

	toy1[0].Name = "c7test"
	toy1[1].Name = "bertest"

	multiUser[0].Toys = toy
	multiUser[1].Toys = toy1
	result = DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&multiUser[0])
	if result.Error != nil {
		t.Errorf("Failed, reupdatesMultiUser[0] error: %v", result.Error)
	}

	// time.Sleep(time.Second * 30)
	result = DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&multiUser[1])
	if result.Error != nil {
		t.Errorf("Failed, reupdatesMultiUser[1] error: %v\n", result.Error)
	}

	//result = DB.Model(&user).Select(clause.Associations).Where("name = ?", "test sister").Delete(&user.ID, 2)
	// result = DB.Model(&user).Select(clause.Associations).Delete(&user.ID, 2)
	result = DB.Model(&user).Select("Toys").Delete(&user.ID, 2)
	if result.Error != nil {
		t.Errorf("Failed, deleteMultiUser[1] error: %v", result.Error)
	}
}
