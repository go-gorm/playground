package main

import (
	"testing"

	"github.com/gofrs/uuid"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	uuid1, _ := uuid.NewV4()
	user := User{Name: "jinzhu", UserID: uuid1}
	uuid2, _ := uuid.NewV4()
	user2 := User{Name: "philicious", UserID: uuid2}

	DB.Create(&user)
	DB.Create(&user2)

	var resultGood []User
	_ = DB.Debug().Where(&User{UserID: uuid1}).Find(&resultGood)
	// resulting query: SELECT * FROM `users` WHERE `users`.`user_id` = "a543742b-9fe4-49ec-9226-e980be5669a8" AND `users`.`deleted_at` IS NULL

	if len(resultGood) > 1 {
		t.Errorf("should only find a single user")
	}

	var resultBad []User
	uuidNull, _ := uuid.FromString("00000000-0000-0000-0000-000000000000") // same as using uuid.Nil or uuid.FromStringOrNil(uuidString)
	_ = DB.Debug().Where(&User{UserID: uuidNull}).Find(&resultBad)
	// resulting query: SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL

	if len(resultBad) > 0 {
		t.Errorf("should NOT find a user")
	}
}
