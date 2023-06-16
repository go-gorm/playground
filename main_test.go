package main

import (
	"gorm.io/gorm"
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

	if err := DB.AutoMigrate(Value{}); err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if err := DB.AutoMigrate(ValueDep{}); err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	// save predefined value to db
	value := Value{}
	if err := DB.Save(&value).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	valueDep := ValueDep{ValueID: value.ID, Name: "some-name", Params: Params{"foo": "bar"}}
	if err := DB.Save(&valueDep).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	// ensure that we can read it
	var value2 Value
	if err := DB.Preload("Deps").First(&value2, value.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	// ensure that the name in db is some-name
	if value2.Deps[0].Name != "some-name" {
		t.Error("Failed to read predefined name value")
	}

	// ensure that the params in db has foo:bar
	if value2.Deps[0].Params["foo"] != "bar" {
		t.Error("Failed to read predefined foo value")
	}

	// update name, foo and save it to db
	value2.Deps[0].Name = "new-name"
	value2.Deps[0].Params["foo"] = "new-bar"
	if err := DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(value2).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	// read again, expect updated values
	var value3 Value
	if err := DB.Preload("Deps").First(&value3, value.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	// ensure simple string was updated
	if value3.Deps[0].Name != "new-name" {
		t.Errorf("Failed to save name, current state is: %v", value3.Deps[0].Name)
	}

	// see that jsonb params wasn't updated
	if value3.Deps[0].Params["foo"] != "new-bar" {
		t.Errorf("Failed to save params, current state is: %v", value3.Deps[0].Params)
	}
}
