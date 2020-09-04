package main

import (
	"testing"
	//"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	DB := DB.Set("Test", "1")
	toy := Toy{
		Name:      "test",
		OwnerID:   "john",
		OwnerType: "humanoid",
	}
	DB.Create(&toy)
	// when we do this instead it works:
	// DB.Session(&gorm.Session{
	// 	WithConditions: true,
	// }).Create(&toy)

	company := Company{
		ID:   1,
		Name: "Evil Corp",
	}
	DB.Create(&company)
}
