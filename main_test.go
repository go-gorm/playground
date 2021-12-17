package main

import (
	"fmt"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{}
	DB.Create(&user)

	pet := Pet{UserID: &user.ID, User: &user}
	DB.Create(&pet)

	user.FavPet = &pet
	user.FavPetID = pet.ID
	DB.Save(&user)

	// This point is never reached as the above code hangs...
	fmt.Println("Wohoo, you got here!!!")
}
