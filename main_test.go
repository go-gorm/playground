package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestInnerJoin(t *testing.T) {
	manager := User{
		Name: "mom",
		Toys: []Toy{
			{
				Name: "t1",
			},
			{
				Name: "t2",
			},
		},
	}
	DB.Create(&manager)
	var result User
	if err := DB.Model(&User{}).Where("users.name = ?", "mom").InnerJoins("Toys", &Toy{Name: "t1"}).Take(&result).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	//if err := DB.Model(&User{}).Where("users.name = ?", "mom").InnerJoins("Toys").Where("toys.name = ?", "t1").Take(&result).Error; err != nil {
	//              t.Errorf("Failed, got error: %v", err)
	//      }

}
