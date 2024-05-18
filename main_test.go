package main

import (
	"testing"

	"gorm.io/gorm/clause"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	addedUsers := []UniqueUser{
		{
			Name: "Alice",
			Age:  18,
		},
		{
			Name: "Bob",
			Age:  22,
		},
		{
			Name: "Cindy",
			Age:  25,
		},
	}

	// Drop TABLE IF EXISTS
	DB.Migrator().DropTable(&UniqueUser{})
	DB.AutoMigrate(&UniqueUser{})

	// create users
	DB.Create(&addedUsers)

	// In this case, I use a new slice, to simulate another business scenario
	// where new data is obtained and needs to be updated
	newUsers := []UniqueUser{}
	// one year later
	for _, user := range addedUsers {
		newUsers = append(newUsers, UniqueUser{
			Name: user.Name,
			Age:  user.Age + 1,
		})
	}

	// update users, use on conflict to update all fields
	DB.Clauses(clause.OnConflict{
		// add Columns to compatible with sqlite
		Columns:   []clause.Column{{Name: "name"}},
		UpdateAll: true,
	}).Create(&newUsers)

	// The bug here is that when using OnConflict{UpdateAll: true} for batch insertion,
	// if there is an update, the returned ID is incorrect

	// In this test case,
	// the expected result is that the second time the insertion is executed,
	// it should actually be updated,
	// so the ID returned by the corresponding object should be the same as the first time.
	for i := range newUsers {
		firseCreateUser := addedUsers[i]
		secondCreateUser := newUsers[i]
		if firseCreateUser.ID != secondCreateUser.ID {
			t.Errorf("Expected ID %d, got %d", firseCreateUser.ID, secondCreateUser.ID)
		}
	}
}
