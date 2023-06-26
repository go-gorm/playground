package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	var u User

	DB.Model(&u).Create(map[string]interface{}{
		"Name": "jinzhu", "Age": 18,
	})

	if u.ID != 0 {
		t.Fatalf("should not fill back")
	}
}
