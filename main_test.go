package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type Person struct {
	Id    int64  `gorm:"column:id;primary_key"`
	Code  string `gorm:"column:code"`
	Name  string `gorm:"column:name"`
	SSN   string `gorm:"->"`
	Email string `gorm:"-"`
}

func TestGORM(t *testing.T) {

	if err := DB.AutoMigrate(Person{}); err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	person := &Person{
		Id:   1,
		Code: "JD",
		Name: "John Doe",
	}

	if err := DB.Create(person).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	person1 := &Person{
		Code: "JD",
		SSN:  "012-74-3045",
	}

	if DB.Where(person1).Find(person1).RowsAffected == 0 {
		t.Error("Should SSN be ignored on select?")
	}

	person2 := &Person{
		Code:  "JD",
		Email: "john@doe.com",
	}

	if err := DB.Where(person2).Find(person2).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

}
