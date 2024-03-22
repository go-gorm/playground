package main

import (
	"fmt"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: mysql

func TestGORM(t *testing.T) {

	//type User struct {
	//	gorm.Model
	//	Name string `gorm:"uniqueIndex;size:100" json:"name"`
	//}
	user1 := &User{
		Name: "user1",
	}
	user2 := &User{
		Name: "user2",
	}
	users1 := []*User{user1, user2}
	DB.Save(users1)
	fmt.Println(user1.ID, user2.ID)

	user2Copy := &User{
		Name: "user2",
	}
	user3 := &User{
		Name: "user3",
	}
	users2 := []*User{user2Copy, user3}
	DB.Save(users2)
	fmt.Println(user2Copy.ID, user3.ID)

	if user2Copy.ID != user2.ID {
		t.Errorf("Failed, got error: %v", "err")
	}
}
