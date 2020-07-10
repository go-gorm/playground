package main

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/clause"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestClear(t *testing.T) {
	user := User{Name: "jinzhu", Pets: []Pet{
		{
			Name: "Pet1",
		},
		{
			Name: "Pet2",
		},
		{
			Name: "Pet3",
		},
	}}

	DB.Create(&user)

	var res1 User
	if err := DB.Preload(clause.Associations).First(&res1, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	assert.Len(t, res1.Pets, 3)

	if err := DB.Model(&res1).Association("Pets").Clear(); err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	var pets []Pet
	if err := DB.Find(&pets).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	assert.Len(t, pets, 0)
}

func TestResave(t *testing.T) {
	user := User{Name: "jinzhu", Pets: []Pet{
		{
			Name: "Pet1",
		},
		{
			Name: "Pet2",
		},
		{
			Name: "Pet3",
		},
	}}

	DB.Create(&user)

	var res1 User
	if err := DB.Preload(clause.Associations).First(&res1, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	assert.Len(t, res1.Pets, 3)

	res1.Pets = []Pet{} // CLEAR ASSOCIATION

	if err := DB.Save(&res1).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	var pets []Pet
	if err := DB.Find(&pets).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	assert.Len(t, pets, 0)
}
