package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	if err := DB.Create(&user).Error; err != nil {
		t.Fatal(err)
	}

	var st2 User
	if err := DB.First(&st2, 1).Error; err != nil {
		t.Fatal(err)
	}

	if st2.ID != 1 {
		t.Fatalf("expected ID 1, got %d", st2.ID)
	}

	st2.Name = "new User"
	if err := DB.Save(&st2).Error; err != nil {
		t.Fatal(err)
	}

	var st3 User
	if err := DB.First(&st3, 1).Error; err != nil {
		t.Fatal(err)
	}
}
