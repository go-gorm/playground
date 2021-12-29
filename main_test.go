package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	// gen a bunch of test User obj
	users := genUsers(11)
	DB.Create(&users)

	// test run for Find fuction on map obj (Find to map)
	var mappedObj map[string]interface{}
	if err := DB.Table("users").Find(&mappedObj).Error; err != nil {
		t.Errorf("Find to map failed, got error: %v", err)
	}

	// fail case for FindInBatches on map obj
	var mappedObjs []map[string]interface{}
	if err := DB.Table("users").FindInBatches(&mappedObjs, 10, func(tx *gorm.DB, batch int) error {
		return nil
	}).Error; err != nil {
		t.Errorf("FindInBatches to map failed, got error: %v", err)
	}
}

func genUsers(n int) []User {
	var users []User
	for i := 0; i < n; i++ {
		u := User{Name: "test"}
		users = append(users, u)
	}
	return users
}
