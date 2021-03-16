package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type (
	Customer struct {
		gorm.Model
		Name      string
		ManagerID *uint
		Manager   *Customer
	}
	CustomerWrapper struct {
		Customer
	}
	Employee struct {
		gorm.Model
		Name       string
		ManagerRef *uint
		Manager    *Employee `gorm:"foreignKey:ManagerRef"`
	}
	EmployeeWrapper struct {
		Employee
	}
)

func (CustomerWrapper) TableName() string {
	return "customer"
}
func (EmployeeWrapper) TableName() string {
	return "employees"
}

func TestGORM(t *testing.T) {
	err := DB.AutoMigrate(new(User), new(Employee))
	if err != nil {
		t.Fatal(err)
	}

	if true {
		manager := &Employee{
			Name: "Boss",
		}
		DB.Create(manager)

		/* schema/relationship.go:guessRelation(relation *Relationship, field *Field, cgl guessLevel)
		if gl == guessGuess {
			if field.Schema == relation.FieldSchema {
				gl = guessBelongs <<
			} else {
				gl = guessHas
			}
		}
		*/

		user1 := &Employee{
			Name:    "User1",
			Manager: manager,
		}
		DB.Create(user1)

		/* schema/relationship.go:guessRelation(relation *Relationship, field *Field, cgl guessLevel)
		if gl == guessGuess {
			if field.Schema == relation.FieldSchema {
				gl = guessBelongs
			} else {
				gl = guessHas <<
			}
		}
		*/

		user2 := &Employee{
			Name:    "User2",
			Manager: manager,
		}
		DB.Create(&EmployeeWrapper{Employee: *user2})

		var m Employee

		DB.First(&m)

		if m.ManagerRef != nil {
			t.Fatal("manager must not have manager")
		}
	} else {
		manager := &Customer{
			Name: "boss",
		}
		DB.Create(manager)

		/* schema/relationship.go:guessRelation(relation *Relationship, field *Field, cgl guessLevel)
		if gl == guessGuess {
			if field.Schema == relation.FieldSchema {
				gl = guessBelongs <<
			} else {
				gl = guessHas
			}
		}
		*/
		user1 := &Customer{
			Name:    "user1",
			Manager: manager,
		}
		DB.Create(user1)

		/* schema/relationship.go:guessRelation(relation *Relationship, field *Field, cgl guessLevel)
		if gl == guessGuess {
			if field.Schema == relation.FieldSchema {
				gl = guessBelongs <<
			} else {
				gl = guessHas
			}
		}
		*/
		user2 := &Customer{
			Name:    "user2",
			Manager: manager,
		}
		DB.Create(&CustomerWrapper{Customer: *user2})

		var m Customer

		DB.First(&m)

		if m.ManagerID != nil {
			t.Fatal("manager must not have manager")
		}
	}

}
