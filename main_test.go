package main

import (
    "fmt"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := []User {{Name: "jinzhu"}, {Name: "unknown"}}

	DB.Create(&user)

	var result []User
	if err := DB.Find(&result).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	} else {
        for i := range user {
            if user[i].ID == result[i].ID {
            fmt.Println("Id matched for user: ", user[i].Name, user[i].ID)
            } else {
                t.Errorf("Id failed for user: %s. Expected ID of %d but got %d\n", user[i].Name, user[i].ID, result[i].ID)
            }
        }
    }
}
