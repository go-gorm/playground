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
		t.Fatalf("Failed, got error: %v", err)
	}

	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("The code did not panic")
			return
		}

		t.Logf("Test paniced as expected: %v", r)
	}()

	// Statement is still set for some reason
	t.Logf("Statement Table: %s", DB.Statement.Table)

	if err := DB.Save(&Language{Code: "Code"}).Error; err == nil {
		t.Fatal("We expected this to fail")
	}
}
