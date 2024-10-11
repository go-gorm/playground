package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// GORM_GEN_REPO: https://github.com/go-gorm/gen.git
// GORM_GEN_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

// When you Assing an attribute and call FirstOrInit, it gets assigned to the result
func TestGORM_Succeeds(t *testing.T) {
	user := User{Name: "jinzhu"}
	manager := User{Name: "haoran"}

	DB.Create(&user)
	DB.Create(&manager)

	var result User
	if err := DB.Where(&User{Name: user.Name}).Assign(&User{Manager: &manager}).FirstOrInit(&result).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	if result.Manager == nil {
		t.Errorf("Failed, user manager is expected")
	}
}

// But hen you Assing an attribute and call FirstOrCreate, it doesn't get assigned to the result
func TestGORM_Fails(t *testing.T) {
	user := User{Name: "jinzhu"}
	manager := User{Name: "haoran"}

	DB.Create(&user)
	DB.Create(&manager)

	var result User
	if err := DB.Where(&User{Name: user.Name}).Assign(&User{Manager: &manager}).FirstOrCreate(&result).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	if result.Manager == nil {
		t.Errorf("Failed, user manager is expected")
	}
}
