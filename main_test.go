package main

import (
	"strings"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu", Nickname: "jinzhu"}
	user2 := User{Name: "john", Nickname: "jinzhu"} // this should fail
	user3 := User{Name: "jane", Nickname: "jinzhu"} // this should fail too

	DB.Create(&user)
	if tx := DB.Create(&user2); tx.Error != nil && strings.Contains(tx.Error.Error(), "UNIQUE constraint failed") {
		t.Log("Got the error", tx.Error) // as expected: UNIQUE constraint failed
	} else {
		t.Error("Did not get error with unique constraint", tx.Error)
	}
	BeforeCreate1 = true  // add first OnConflict clause for `nickname`
	if tx := DB.Create(&user2); tx.Error != nil { // this should not fail
		t.Error("Still got error", tx.Error)
	}
	BeforeCreate2 = true  // add second OnConflict clause for `id`
	// (in fact it overwrites the first one, resulting in no OnConflict clauses for nickname
	if tx := DB.Create(&user3); tx.Error != nil { // this should not fail
		t.Errorf("Getting error because of overwritten clause: %v\nNotice the missing first ON CONFLICT clause in the query above!", tx.Error)
	}

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
