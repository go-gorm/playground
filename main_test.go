package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	u := &User{Name: "george"}
	DB.Create(&u)

	DB.Create(&User{Name: "jinzhu", Friends: []*User{u}})

	george := &User{}
	DB.Model(&User{}).Where("name = ?", "george").Find(george)
	if err := george.UUID.String() != u.UUID.String(); err {
		t.Error("Failed, u.UUID has changed but the the one in the DB is the same")
	}

}
