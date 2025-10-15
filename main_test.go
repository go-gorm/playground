package main

import (
	"testing"

	"gorm.io/playground/models"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	acc := models.Account{Number: "foo"}
	DB.Create(&acc)
	acc2 := models.Account{Number: "baz"}
	DB.Create(&acc2)

	user := models.User{AccountId: &acc.ID, Name: "Bar"}

	DB.Create(&user)

	var user2 = models.User{}
	DB.Preload("Account").Where(models.User{Name: "Bar"}).FirstOrInit(&user2)

	id_backup := acc2.ID

	// This should update the AccountId to acc2.ID.
	// However, when a relation was loaded before, this will actually reset back to acc.ID
	// When using nullable fields like here, it will also overwrite the ID in acc2
	user2.AccountId = &acc2.ID
	DB.Save(&user2)

	if  *user2.AccountId != id_backup {
		t.Errorf("Account ID not properly saved: %v", *user2.AccountId)
	}
}
