package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "tom"}
	err := DB.Create(&user).Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	pets := [...]*Pet{
		&Pet{
			Name:   "fluffy",
			UserID: &user.ID,
		},
		&Pet{
			Name:   "lucky",
			UserID: &user.ID,
		},
	}

	for _, pet := range pets {
		err := DB.Create(pet).Error
		if err != nil {
			t.Errorf("Failed, got error: %v", err)
		}
	}

	var u User
	err = DB.Joins("Pets").Where("users.id = ?", user.ID).Find(&u).Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

}
