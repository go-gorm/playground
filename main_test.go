package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result User

	// create the SQL query to display a bug
	query := "SELECT * FROM users WHERE name = ?"

	// run the good raw SQL query to display the query is valid
	goodResult := DB.Raw(query, "jinzhu").Scan(&result)
	// should return no error since user 'jinzhu' exists
	if err := goodResult.Error; err != nil {
		t.Errorf("Gorm Raw() Failed, got: %v, want: %v", err, nil)
	}

	// run the bad raw SQL query to display the bug
	badResult := DB.Raw(query, "foobar").Scan(&result)
	// should return gorm.ErrRecordNotFound since user 'foobar' does not exist
	if err := badResult.Error; err == nil {
		t.Errorf("Gorm Raw() Failed, got: %v, want: %v", err, gorm.ErrRecordNotFound)
	}
}
