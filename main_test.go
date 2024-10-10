package main

import (
	"testing"

	"gopkg.in/guregu/null.v3"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {

	user := User{Name: "jinzhu", OptionalID: null.StringFrom("this_is_an_optional_id")}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	// Update the user's name and optionalID (to be NULL)
	toUpdate := user
	toUpdate.OptionalID = null.String{}
	toUpdate.Name = "Steve"

	err := DB.Save(&toUpdate).Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	// Refetch the updated user into &user
	if err := DB.Model(user).
		First(&user).
		Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	// Name is updated
	if user.Name != "Steve" {
		t.Errorf("Updated name was not fetched")
	}

	// OptionalID isn't updated (should be invalid)
	if user.OptionalID.Valid {
		t.Errorf("Updated optional ID was not fetched")
	}
}
