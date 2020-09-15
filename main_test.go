package main

import (
	"gorm.io/gorm/clause"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type Asset struct {
	ID         int
	Kind       string
	Value      float32
	BusinessID int
}

type Business struct {
	ID       int
	Name     string
	Assets   []Asset `gorm:"foreignkey:BusinessID;"`
	PersonID int
}

type Person struct {
	ID       int
	Name     string
	Business Business   `gorm:"foreignkey:PersonID;"`
}

func TestGORM(t *testing.T) {

	err := DB.AutoMigrate(&Person{}, &Business{}, &Asset{})
	if err != nil {
		panic(err)
	}

	if err = DB.Create(&Person{
		Name: "Jinzhu",
		Business: Business{
			Name: "GORM",
			Assets: []Asset{
				{
					Kind: "Cash",
					Value: 10000,
				},
			},
		},
	}).Error; err != nil {
		panic(err)
	}

	people := make([]Person, 0)

	// If I do the following, it works:
	// db.Preload("Company").Preload("Company.Address").Preload("Company.Assets")
	// But not if I do this:
	err = DB.Preload(clause.Associations).Find(&people).Error
	if err != nil {
		panic(err)
	}
	if people[0].Business.Assets == nil {
		panic("despite having assets, they are not preloaded, and are nil")
	}
}
