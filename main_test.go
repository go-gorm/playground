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
		t.Errorf("Failed, got error: %v", err)
	}

	test := Test{Data: []float64{8, 4, 2, 1, 0.5}}

	err := DB.Create(&test).Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	var result2 Test
	if err := DB.First(&result2).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
