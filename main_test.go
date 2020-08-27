package main

import (
	"testing"
	"fmt"
)

type User struct {
	ID           int32 `gorm:"primaryKey"`
	Name         string
	CreatedBy    *int32
	Creator      *User `gorm:"foreignKey:CreatedBy;references:ID"`
}

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	u1 := User{ID: 1, CreatedBy: nil}
	DB.Create(&u1)
	
	u2 := User{ID: 2, CreatedBy: &u1.ID}
	DB.Create(&u2)

	u1Found := &User{}
	DB.Preload("Creator").Find(u1Found, 1)
	fmt.Println(u1Found) // CreatedBy is nil, but Creator is u2

	u2Found := &User{}
	DB.Preload("Creator").Find(u2Found, 2)
	fmt.Println(u2Found) // CreatedBy is 1, but Creator is nil
}
