package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver


func TestGORM(t *testing.T) {
	// Prepare test data
	t1, _ := time.Parse(time.RFC3339, "2020-09-29T00:00:00Z")
	t2, _ := time.Parse(time.RFC3339, "2020-10-30T00:00:00Z")

	// Fill DB with test data
	DB.Create(&User{Name: "Alice", Birthday: &t1})
	DB.Create(&User{Name: "Alice", Birthday: &t2})
	DB.Create(&User{Name: "Bob", Birthday: &t1})
	DB.Create(&User{Name: "Bob", Birthday: &t2})
	DB.Create(&User{Name: "Carla", Birthday: &t1})
	DB.Create(&User{Name: "Carla", Birthday: &t2})

	// Query DB to reproduce bug
	var result []User

	query := DB.Model(&User{})

	// We add a where clause with OR statement
	query = query.Where("1=1").Or("name = ?", "Bob")

	// We add another where clause
	query = query.Where("Birthday <= ?", t1)

	// At this stage, the SQL behind the stage should be where (name = 'Alice' or name = 'Bob') and Birthday <= '2020-09-29T00:00:00Z'

	if err := query.Find(&result).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	for _, user := range result {
		fmt.Println(fmt.Sprintf("User: %s, dob: %s", user.Name, user.Birthday))
	}
}

