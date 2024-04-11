package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	c := C{}
	DB.Create(&c)
	b := B{C: c}
	DB.Create(&b)
	a := A{B: b}
	DB.Create(&a)
	base := Base{A: &a}
	DB.Create(&base)

	var res Base
	// Nested query that results in unsupported data
	if err := DB.Table("bases").Joins("A").Joins("A.B").Preload("A.B.C").First(&res).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
