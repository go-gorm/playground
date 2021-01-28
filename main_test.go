package main

import (
	"fmt"
	"gorm.io/gorm/clause"
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
}

func TestAppendNested(t *testing.T){
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	// Add first pet, all good
	pet1 := Pet{
		Name:   "Pet1",
		IsDog:  false,
	}

	err := DB.Model(&user).Association("Pets").Append(&pet1)
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	var result User
	if err := DB.Preload(clause.Associations).First(&result, user.ID).Error; err != nil {
			t.Errorf("Failed, got error: %v", err)
	}

	if len(result.Pets) != 1 {
		t.Errorf("Should be only 1 pet, but is: %v", result.Pets)
	}

	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	// Add second pet, all good
	pet2 := Pet{
		Name:   "Pet2",
		IsDog:  false,
	}
	err = DB.Model(&user).Association("Pets").Append(&pet2)
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if err := DB.Preload(clause.Associations).First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if len(result.Pets) != 2 {
		t.Errorf("Should be only 2 pet, but is: %v", result.Pets)
	}

	// Add third pet, which is the same as first and should throw something like a unique constraint exception
	pet3 := Pet{
		Name:   "Pet1",
		IsDog:  false,
	}

	expectedErr := DB.Model(&user).Association("Pets").Append(&pet3)

	if err := DB.Preload(clause.Associations).First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if len(result.Pets) != 2 {
		t.Errorf("Should be only 2 pet, but is: %v", result.Pets)
	} else {
		fmt.Println("We have only 2 pets. Which is the expected result. But i would an error already when trying to append")
	}

	if expectedErr == nil {
		t.Errorf("I would expect error here")
	}

}
