package main

import (
	"strconv"
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

func TestUpdateWithLimit(t *testing.T) {
	defer DB.Delete(User{}, "name LIKE ?", "user%")

	for i := 0; i < 10; i++ {
		user := User{Name: "user" + strconv.Itoa(i), Age: uint(i)}
		DB.Create(&user)
	}

	tx := DB.Model(&User{}).Select("age").Where("name LIKE ?", "user%").Order("name ASC").Limit(5).Updates(&User{Age: 100})

	if tx.Error != nil {
		t.Errorf("Failed, got error: %v", tx.Error)
	}

	if tx.RowsAffected > 5 {
		t.Errorf("Failed, update %v rows", tx.RowsAffected)
	}
}
