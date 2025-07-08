package main

import (
	"fmt"
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

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu", IsDeleted: false}
	DB = DB.Debug()

	DB.Create(&user)

	var user1 User
	basicCondition := map[string]string{
		fmt.Sprintf("%s.is_deleted", "users"): "false",
	}
	err := DB.Table("users").Where(basicCondition).Find(&user1)
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	fmt.Printf("User: %v\n", user)
}
