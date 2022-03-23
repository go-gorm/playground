package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	var pets []*Pet

	err := DB.
		Joins("INNER JOIN users ON users.id = pets.user_id").
		Where(&Pet{Name: "a"}).
		Where(&User{Age: 10}).
		Find(&pets).Error

	if err != nil {
		t.Fatalf("err: %v", err)
	}
}
