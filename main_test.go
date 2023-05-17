package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	if err := DB.AutoMigrate(&Thing1{}, &Thing2{}, &Thing3{}); err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	DB.Create(&Thing1{Name: "Thing 1", One: 1})
	DB.Create(&Thing2{Name: "Thing 2", Two: 2})
	DB.Create(&Thing3{Name: "Thing 3", Three: 3})

	var result Composite
	err := DB.Table("thing1").
		// worked
		// Joins("Thing2").
		// not worked
		// Joins("LEFT JOIN thing2 Thing2 ON thing1.id1=thing2.id2").
		// not worked
		Joins("LEFT JOIN thing2 Thing2 ON thing1.id1=thing2.id2 AND thing2.two=2").
		Joins("Thing3").
		Find(&result).Error
	if err != nil {
		t.Fatal(err)
	}
	if result.Thing2.Name != "Thing 2" || result.Thing2.Two != 2 {
		t.Error("embed Thing2 not right")
	}
	if result.Thing3.Name != "Thing 3" || result.Thing3.Three != 3 {
		t.Error("embed Thing3 not right")
	}
}
