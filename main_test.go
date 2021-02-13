package main

import (
	"fmt"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}

func TestSubQueryRaw(t *testing.T) {
	subQ := DB.Raw(`SELECT users.id FROM users WHERE id IS NOT NULL`)

	q := DB.Table("users").Where("id IN (?)", subQ)

	var users []*User
	err := q.Find(&users).Error

	if err == nil {
		t.Fail()
	}

	fmt.Printf("%v", err)
	//no such table: ---
	// Printed SQL is:
	//  SELECT * FROM `users` WHERE id IN (SELECT * FROM ``) AND `users`.`deleted_at` IS NULL
}

func TestSubQueryBuilder(t *testing.T) {
	subQ := DB.Table("users").Select("id").Where("id IS NOT NULL")

	q := DB.Table("users").Where("id IN (?)", subQ)

	var users []*User
	err := q.Find(&users).Error

	if err != nil {
		fmt.Printf("%v", err)
		t.Fail()
	}
}
