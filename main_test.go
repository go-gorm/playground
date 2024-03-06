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

	// this works fine.
	users := []User{}
	company := Company{}

	err := DB.Raw("select * from companies limit 1").Scan(&company).Error
	if err != nil {
		t.Errorf("Failed, got error: %+v", err)
	}

	// this does not panic and works as expected.
	err = DB.Raw("select * from users limit 2").Scan(&users).Error
	if err != nil {
		t.Errorf("Failed, got error: %+v", err)
	}

	// this panics!
	err = DB.Scopes(func(d *gorm.DB) *gorm.DB {
		// If Block A is removed, the panic goes away.
		//
		// BLOCK A
		a := Account{}
		if err := d.Raw("select * from accounts limit 1").Scan(&a).Error; err != nil {
			return d
		}
		// END BLOCK

		return d.Raw("select * from users limit 2")
	}).Scan(&users).Error

	if err != nil {
		t.Errorf("Failed, got error: %+v", err)
	}
}
