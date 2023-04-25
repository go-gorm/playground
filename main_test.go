package main

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type ProfileState byte

const (
	ProfileStateDraft ProfileState = iota
	ProfileStatePublished
)

var (
	profileStates        []ProfileState
	profileStateToString = map[ProfileState]string{
		ProfileStateDraft:     "draft",
		ProfileStatePublished: "published",
	}
	profileStateToID = map[string]ProfileState{
		"draft":     ProfileStateDraft,
		"published": ProfileStatePublished,
	}
)

func (ct ProfileState) String() string {
	return profileStateToString[ct]
}

//Scan gorm scans the value in db and parse it to enum
func (ct *ProfileState) Scan(value interface{}) error {
	var exists bool

	*ct, exists = profileStateToID[value.(string)]
	if !exists {
		return errors.New("ProfileState scan: invalid value")
	}

	return nil
}

//Value gorm gets the value to store it in the db
func (rt ProfileState) Value() (driver.Value, error) {
	return profileStateToString[rt], nil
}

type MyUser struct {
	*User
	Status ProfileState `gorm:"type:citext;"`
}

func TestGORM(t *testing.T) {
	DB.AutoMigrate(&MyUser{})

	user := MyUser{
		User: &User{
			Name: "jinzhu",
		},
		Status: ProfileStateDraft,
	}

	DB.Save(&user)

	var result MyUser
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	fmt.Println(result)

	user.Status = ProfileStatePublished
	DB.Save(&user)

	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	fmt.Println(result)
}
