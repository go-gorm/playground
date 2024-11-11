package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// GORM_GEN_REPO: https://github.com/go-gorm/gen.git
// GORM_GEN_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}

func TestGORM_failure(t *testing.T) {
	value := &Entity{Name: "value"}

	DB.Create(value)

	DB.Save(value)
}
