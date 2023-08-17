package main

import (
	"testing"

	"gorm.io/playground/models"
	"gorm.io/playground/query"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	query.SetDefault(DB)
	user := models.User{
		Name: "jinzhu",
		Pets: []*models.Pet{
			{
				Name: "my-pet",
				Toy:  models.Toy{},
			},
		},
	}

	// Test passes when query.Q is used
	tx := query.Q

	// But when using transactions it stops at `Replace`
	// tx := query.Q.Begin()
	// defer tx.Rollback()

	if err := tx.User.Save(&user); err != nil {
		t.Fatal(err)
	}
	if err := tx.User.Pets.Model(&user).Replace(user.Pets...); err != nil {
		t.Fatal(err)
	}
}
