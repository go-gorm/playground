package main

import (
	"encoding/json"
	"log"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver
func TestGORM(t *testing.T) {
	str := `{"userName" = "test", "age" = 18 }`

	bytes, err := json.Marshal(str)
	if err != nil {
		log.Printf("Marshal failed,err = %+v", err)
	}

	user := User{
		ID:          1,
		Description: bytes,
	}

	if err := DB.Create(&user).Error; err != nil {
		log.Fatal(err)
	}

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	log.Println(string(user.Description))
}
