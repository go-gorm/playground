package main

import (
	"fmt"
	"log"
	"testing"

	"gorm.io/gorm"
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

func TestShouldInsert(t *testing.T) {
	if err := DB.AutoMigrate(Artist{}); err != nil {
		t.Fatalf("error migrating releases table %s", err)
	}

	tx := DB.Begin().Debug()
	defer tx.Rollback()

	// The expected result is that an artist should be inserted, but for some
	// reason gorm tries to update, which fails due to a lack of where-clause.
	if err := doArtist(tx, "new artist name"); err != nil {
		t.Fatalf("error %s", err)
	}
}

func doArtist(tx *gorm.DB, name string) (err error) {
	artist := &Artist{}
	err = tx.Where(Artist{Name: name}).
		Attrs(Artist{Name: name}).
		FirstOrInit(&artist).Error
	if err != nil {
		return fmt.Errorf("initializing artist %s", err)
	}
	log.Printf("<<<<>>>> the new artist's id is %d <<<<>>>>", artist.ID)
	err = tx.Save(&artist).Error
	if err != nil {
		return fmt.Errorf("saving artist %s", err)
	}

	return nil
}

type Artist struct {
	gorm.Model
	Name string
}
