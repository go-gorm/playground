package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu", Gender: 1, Cond: "test"}
	user2 := User{Name: "jinzhu", Gender: 2, Cond: "test"}
	DB.Create(&user)
	DB.Create(&user2)
	value := []User{}
	err := DB.Where("cond = ?", "test").FindInBatches(&value, 1, func(tx *gorm.DB, batch int) error {
		// Do something...
		// error occurs in FindInBatches
		return nil
	}).Error
	if err != nil {
		t.Errorf("%+v", err)
	}
}
