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

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	sub1 := DB.Model(&User{}).Select("name", "age").Where("age = ?", 30)
	sub2 := DB.Table("(?) as u2", sub1).Select("name", "age").Where("age = ?", 18)
	DB.Table("(?) as u", sub2).Where("name = ?", 20).Find(&User{})
}
