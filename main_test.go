package main

import (
	"reflect"
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

	if ok := reflect.DeepEqual(&user, &result); !ok {
		t.Errorf("Failed, object not match: \ncreate: %+v\nget: %+v", user, result)
	}
}
