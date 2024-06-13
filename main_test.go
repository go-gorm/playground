package main

import (
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

func TestPluckWithNullableColumns(t *testing.T) {
	var ids []*int64
	if err := DB.Model(&User{}).Where("name = 'dz'").Pluck("MAX(id)", &ids).Error; err != nil {
		t.Errorf("got error on calling Pluck(): %v", err)
	}
	// ids should be []*int64{nil}
	if len(ids) != 1 || ids[0] != nil {
		t.Errorf("wrong result: %#v", ids)
		if len(ids) > 0 && ids[0] != nil {
			t.Errorf("first element's value: %#v", *ids[0])
		}
	}
}
