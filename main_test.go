package main

import (
	"fmt"
	"gorm.io/gorm"
	"testing"
	"time"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: mysql

type Person struct {
	ID        int
	Name      string
	Addresses []Address `gorm:"many2many:person_addresses;"`
}

type Address struct {
	ID   uint
	Name string
}

type PersonAddress struct {
	PersonID  int
	AddressID int
	Home      bool
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (p *PersonAddress) AfterCreate(db *gorm.DB) error {
	db.Model(p).Where(p).Update("home", true)
	return nil
}

func TestGORM(t *testing.T) {

	DB.AutoMigrate(&Address{})
	DB.AutoMigrate(&Person{})
	DB.AutoMigrate(&PersonAddress{})

	err := DB.SetupJoinTable(&Person{}, "Addresses", &PersonAddress{})

	if err != nil {
		fmt.Println(err.Error())
	}

	DB.Debug().Create(&Person{
		ID:   1,
		Name: "Joe",
		Addresses: []Address{{
			ID:   1,
			Name: "First Address",
		}},
	})
}
