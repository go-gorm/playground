package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var names []string
	DB.Raw(`select name from users`).Pluck("name", &names)

	if len(names) != 1 {
		t.Errorf("Failed, Pluck without Table() didn't return any results")
	}

	DB.Table("blag").Raw(`select name from users`).Pluck("name", &names)
	if len(names) != 1 {
		t.Errorf("Failed, Pluck without Table() didn't return any results")
	}
	// This succeeds, even though "blag" isn't even a real table.
}
