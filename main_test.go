package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{}
	DB.Create(&user)

	pet1 := Pet{Name: "pet1", UserID: user.ID}
	pet2 := Pet{Name: "pet2", UserID: user.ID}

	DB.Create(&pet1)
	DB.Create(&pet2)

	toy1 := Toy{Name: "toy1", PetID: pet1.ID}
	toy2 := Toy{Name: "toy2", PetID: pet1.ID}
	toy3 := Toy{Name: "toy3", PetID: pet2.ID}
	toy4 := Toy{Name: "toy4", PetID: pet2.ID}

	DB.Create(&toy1)
	DB.Create(&toy2)
	DB.Create(&toy3)
	DB.Create(&toy4)

	pet1.FavToyID = toy1.ID
	DB.Save(&pet1)

	pet2.FavToyID = toy4.ID
	DB.Save(&pet2)

	userFromDB := &User{}
	result := DB.
		Preload("Pets").
		Preload("Pets.FavToy").
		First(userFromDB)

	if result.Error != nil {
		t.Error("Got unexpected result from query: ", result.Error)
	}

	if len(userFromDB.Pets) != 2 {
		t.Errorf("Expected 2 pets but got %v", len(userFromDB.Pets))
	}

	var pet1FromDB *Pet
	var pet2FromDB *Pet

	if userFromDB.Pets[0].Name == pet1.Name {
		pet1FromDB = userFromDB.Pets[0]
		pet2FromDB = userFromDB.Pets[1]
	} else {
		pet1FromDB = userFromDB.Pets[1]
		pet2FromDB = userFromDB.Pets[0]
	}

	if pet1FromDB.FavToy.ID != toy1.ID {
		t.Errorf("Expected pet1's fav toy to be '%v' but got '%v'", toy1.Name, pet1FromDB.FavToy.Name)
	}

	if pet2FromDB.FavToy.ID != toy4.ID {
		t.Errorf("Expected pet2's fav toy to be '%v' but got '%v'", toy4.Name, pet2FromDB.FavToy.Name)
	}
}
