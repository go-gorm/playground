package main

import (
	"testing"

	"gorm.io/gorm/clause"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: mysql

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}

func TestConfclictUpdateAll(t *testing.T) {
	p := Product{}
	p.Name = "test"
	p.Age = 10

	DB.Create(&p)

	var result Product
	if err := DB.First(&result, p.Id).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	if result.Name != p.Name {
		t.Error("Name not saved")
	}

	if result.Age != p.Age {
		t.Error("Age not saved")
	}

	db := DB.Clauses(clause.OnConflict{UpdateAll: true})
	p = Product{}
	p.Name = "test2"
	p.Age = 12
	db.Create(&p)
	if err := db.First(&result, p.Id).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	if result.Name != p.Name {
		t.Error("Name not updated")
	}

	if result.Age != p.Age {
		t.Error("Age not updated")
	}
}
