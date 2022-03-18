package main

import (
	"fmt"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {

	add := make([]User, 0)
	for i := 0; i < 100000; i++ {
		user := User{Name: fmt.Sprintf("%d", i)}
		add = append(add, user)
	}

	DB.CreateInBatches(add, 1000)

	fmt.Println("done")

	res := make([]User, 0)

	err := DB.Preload("Friends").Find(&res).Error
	if err != nil {
		t.Error(err)
	}
}
