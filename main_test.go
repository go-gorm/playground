package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu", Toys: []Toy{{TestArray: []int64{1}}}}
	user2 := User{Name: "jorge", Toys: []Toy{{TestArray: []int64{1, 2, 3}}}}

	if err := DB.Create(&user).Error; err != nil {
		t.Errorf("unable to create record: %v", err)
	}

	if err := DB.Create(&user2).Error; err != nil {
		t.Errorf("unable to create record: %v", err)
	}

	var result User
	if err := DB.First(&result, user2.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
