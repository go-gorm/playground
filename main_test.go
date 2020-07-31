package main

import (
	"gorm.io/gorm"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	pet := Pet{
		UserID: &user.ID,
		Name:   "Test pet",
	}
	DB.Create(&pet)
	//var newPet Pet
	var result User
	var count int64

	done := make(chan bool, 1)

	queryDb := DB.Joins("left join pets on pets.user_id = users.id ")
	go countRecords(queryDb, &result, done, &count)

	queryDb.Select("users.name as user_name","pets.name as name").Limit(1000).Find(&result)
	<-done

	if count != 1 {
		t.Errorf(" Total count doesn't match ")
	}
}

func countRecords(db *gorm.DB, anyType interface{}, done chan bool, count *int64) {
	db.Model(anyType).Count(count)
	done <- true
}
