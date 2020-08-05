package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var c int64 = 0

	if err := DB.Table("users").Where("name = ?", "jinzhu").Select("name").Count(&c).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	for i := uint(0); i < 10000; i ++ {
		err := DB.Updates(User{Name: "jinzhu", Age: i}).Error
		if err != nil {
				t.Errorf("gorm errored :%v ", err)
		}
	}
}