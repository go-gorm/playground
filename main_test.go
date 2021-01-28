package main

import (
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

func TestReplace(t *testing.T){
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

	// Add pet3, all good
	pet3 := Pet{
		Name:   "Pet3",
		IsDog:  false,
	}
	err = DB.Model(&user).Association("Pets").Append(&pet3)
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	var result User
	if err := DB.Preload(clause.Associations).First(&result, user.ID).Error; err != nil {
			t.Errorf("Failed, got error: %v", err)
	}

	if len(result.Pets) != 2 {
		t.Errorf("Should be 2 pets, but is: %s", result.Pets)
	}

	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	// Replace with pet2
	pet2 := Pet{
		Name:   "Pet2",
		IsDog:  false,
	}
	var pets []Pet
	pets = append(pets, pet2)
	err = DB.Model(&user).Association("Pets").Replace(&pet2)
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if err := DB.Preload(clause.Associations).First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if result.Pets[0].Name != pet2.Name {
		t.Errorf("Should be %s but is %s", pet2.Name, result.Pets[0].Name)
	}

	if len(result.Pets) != 1 {
		t.Errorf("Should be exactly 1 pet, but is: %+v", &result.Pets)
	}


	// Now replace with adjusted pet2
	pet2.Name = "Adjusted Pet2"
	err = DB.Model(&user).Association("Pets").Replace(&pet2)
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if err := DB.Preload(clause.Associations).First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if result.Pets[0].Name != "Adjusted Pet2" {
		t.Errorf("Should be %s but is %s", "'Adjusted Pet2'", result.Pets[0].Name)
	}

	if len(result.Pets) != 1 {
		t.Errorf("Should be exactly 1 pet, but is: %+v", &result.Pets)
	}


}
