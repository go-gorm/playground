package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	petTypes := []PetType{{Type: "Cat"}, {Type: "Dog"}}
	DB.Create(&petTypes)

	pet := Pet{Name: "Felix", PetType: petTypes[0]}

	DB.Create(&pet)

	var result []Pet
	if err := DB.Joins("PetType").Find(&result, "PetType.Type = ?", "Cat").Error; err != nil {
		t.Errorf("failed, got error: %v", err)
	}
	t.Logf("%+v", result)
}
