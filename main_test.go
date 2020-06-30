package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user1 := User{Name: "jinzhu 1"}
	user2 := User{Name: "jinzhu 2"}
	user3 := User{Name: "jinzhu 3"}
	user4 := User{Name: "jinzhu 4"}

	DB.Create(&user1)
	DB.Create(&user2)
	DB.Create(&user3)
	//user4 not inserted

	if err := DB.Model(&user4).Updates(map[string]interface{}{"name": "test name"}).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	var result []User
	DB.Find(&result)
	t.Logf("Result: %v", result)
}
