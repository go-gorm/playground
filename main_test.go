package main

import (
	"strings"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}
	// expect that it keeps the exact insertion fields order in RETURNING clause
	returningPart := `RETURNING "id","address"`
	tx := DB.Create(&user)
	t.Errorf(tx.Statement.SQL.String())
	if !strings.Contains(tx.Statement.SQL.String(), returningPart) {
		t.Errorf("Expected fields in RETURNING clause preserve insertion order")
	}

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
