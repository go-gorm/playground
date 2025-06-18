package main

import (
	"math/rand"
	"testing"

	"gorm.io/gorm"
)

func (x *User) BeforeCreate(tx *gorm.DB) error {
	if x.ID == 0 {
		x.ID = uint(rand.Uint32())
	}

	return nil
}

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlserver
func TestGORM(t *testing.T) {
	manager := &User{
		Name: "Manager",
	}
	user1 := User{
		Name:    "jinzhu1",
		Manager: manager,
	}
	user2 := User{
		Name:    "jinzhu2",
		Manager: manager,
	}
	user3 := User{
		Name:    "jinzhu3",
		Manager: manager,
	}

	tx := DB.Create([]*User{&user1, &user2, &user3})
	if tx.Error != nil {
		t.Errorf("Failed, got error: %v", tx.Error)
	}
}
