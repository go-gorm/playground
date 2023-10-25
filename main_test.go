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

	p := Pet{Name: "max", UserID: &user.ID}

	DB.Create(&p)

	var result Pet
	if err := DB.First(&result, p.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	err := DB.Migrator().DropColumn(&User{}, "age")
	if err != nil {
		t.FailNow()
	}
}
