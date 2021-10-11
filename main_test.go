package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	companyA := Company{Name: "A"}
	companyB := Company{Name: "B"}
	DB.Create(&companyA)
	DB.Create(&companyB)

	user := User{Name: "jinzhu", CompanyID: &companyB.ID}
	DB.Create(&user)

	query := DB.Model(&User{}).Joins("Company")

	// Bug happens when .Count is called on a query.
	// Removing the below two lines or downgrading to gorm v1.20.12 will make this test pass.
	var total int64
	query.Count(&total)

	var result []User

	if err := query.Find(&result).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	if result[0].Company.ID == 0 {
		t.Errorf("Failed, expected Company to be preloaded")
	}
}
