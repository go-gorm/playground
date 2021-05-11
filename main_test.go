package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	// create some users
	users := []User{
		User{Name: "A", Age: 10},
		User{Name: "B", Age: 18},
		User{Name: "C", Age: 24},
		User{Name: "D", Age: 32},
	}

	DB.Create(&users)

	var results []User
	if err := DB.Find(&results).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if len(users) != len(results) {
		t.Errorf("Failed, expected to create %d users but only %d were created",
			len(users), len(results))
	}

	for _, u := range results {
		if u.DeletedBy != "" {
			t.Errorf("Failed, DeletedBy should be empty for UserID %d", u.ID)
		}
	}

	// delete some users
	users = users[2:4]
	if err := DB.Delete(&User{}, "age > 21").Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if err := DB.Find(&results).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if len(users) != len(results) {
		t.Errorf("Failed, expected to create %d users but only %d were created",
			len(users), len(results))
	}

	for _, u := range results {
		if u.DeletedBy != "gabriel" {
			t.Errorf("Failed, DeletedBy should not be empty for UserID %d", u.ID)
		}
	}
}
