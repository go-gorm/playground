package main

import (
	"testing"

	automigv1 "gorm.io/playground/automig/v1"
	automigv2 "gorm.io/playground/automig/v2"
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

func TestAutoMigrate(t *testing.T) {
	req1 := automigv1.Request{Name: "Request1"}
	req2 := automigv1.Request{Name: "Request1"}

	err := DB.AutoMigrate(automigv1.Request{})
	if err != nil {
		t.Fatalf("unable to automigrate to v1 %v", err)
	}

	err = DB.Create(&req1).Error
	if err != nil {
		t.Fatalf("At least first object should have been created: %v", err)
	}
	err = DB.Create(&req2).Error
	if err == nil {
		t.Fatalf("There should be an unique index violation on Name")
	}

	err = DB.AutoMigrate(automigv2.Request{})
	// At this point the unique index on name should have been deleted
	// and it should be possible to create the req2 object
	if err != nil {
		t.Fatalf("unable to automigrate to v2 %v", err)
	}

	err = DB.Create(&req2).Error
	if err != nil {
		t.Fatalf("req2 should be created after migrating to v2: %v", err)
	}

}
