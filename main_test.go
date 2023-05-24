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

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	// first, demonstrate that the email address is empty in the db.
	if result.EmailNoDefault != "" {
		t.Errorf("expected email to be empty but got: %s", result.Email)
	}
	if result.Email != "" {
		t.Errorf("expected email to be empty but got: %s", result.Email)
	}

	// now, we'll set it to a value before reading it from the db.
	result.Email = "jinzhu@example.com"
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	// shouldn't the email address be empty after the read, since it's not set
	// in the db?
	if result.EmailNoDefault != "" {
		t.Errorf("expected email to be empty but got: %s", result.Email)
	}
	if result.Email != "" {
		t.Errorf("expected email to be empty but got: %s", result.Email)
	}

}
