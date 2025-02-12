package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	// we do not use any default transactions, just explicit ones
	DB.Config.SkipDefaultTransaction = true
	DB.SkipDefaultTransaction = true
	DB.Callback().Create().After("gorm:commit_or_rollback_transaction").Register("example-post-transaction-hook", func(db *gorm.DB) {
		// get the list of all users in the database
		var usersInHook []*User
		err := DB.Find(&usersInHook).Error
		if err != nil {
			t.Fatal(err)
		}
		// we have created a user and are running after the commit or rollback transaction hook and
		// since our transaction was not rolled back, our user should exist
		if len(usersInHook) == 0 {
			t.Fail()
			t.Logf("got %v users after transaction commit", len(usersInHook))
		}
		// as an example, we would like to log the names of all existing users, any time a new user is created
		for _, user := range usersInHook {
			t.Log(user.Name)
		}
	})
	// an example transaction that creates three users
	err := DB.Transaction(func(tx *gorm.DB) error {
		tx.Create(&User{Name: "jinzhu"})
		tx.Create(&User{Name: "Peter"})
		tx.Create(&User{Name: "Bert"})
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
