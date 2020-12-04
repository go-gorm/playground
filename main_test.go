package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	count := 100000
	users := make([]User, count)
	for i := 0; i < count; i++ {
		users = append(users, User{Name: "rbren"})
	}
	DB.Create(users)
}
