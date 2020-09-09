package main

import (
	"testing"
	"gorm.io/datatypes"
	"gorm.io/gorm"

)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: postgres

type UserWithJSON struct {
	gorm.Model
	Name       string
	Attributes datatypes.JSON
}



func TestGORM(t *testing.T) {
	DB.AutoMigrate(&UserWithJSON{})

	err := DB.Create(&UserWithJSON{
		Name:       "json-1",
		Attributes: datatypes.JSON([]byte(`{"name": "jinzhu", "age": 18, "tags": ["tag1", "tag2"], "orgs": {"orga": "orga"}}`)),
	}).Error

	if err != nil {
		t.Errorf("datatypes.json create error")
	}
}
