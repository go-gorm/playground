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

	//Version == 1
	DB.Create(&user)

	//Version == 2
	if err := DB.Model(&user).Select("version", "name", "age").Updates(map[string]interface{}{"Name": "Marck", "Age": 20}).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if result.Version != 2 {
		t.Errorf("Failed, the result should be: 2")
	}

	fmt.Println(result)
}
