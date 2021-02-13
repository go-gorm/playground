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
}

func TestDoubleStatementExect(t *testing.T) {
	var result User
	err := DB.Raw(`
			SELECT * FROM users WHERE id = 0;
			SELECT * FROM users WHERE id = 0;
		`).Scan(&result).Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
		//ERROR: cannot insert multiple commands into a prepared statement (SQLSTATE 42601); ERROR: cannot insert multiple commands into a prepared statement (SQLSTATE 42601)
	}
}
