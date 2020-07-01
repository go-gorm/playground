package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	companyID := 5
	user := User{
		Name: "jinzhu",
		Pets: []*Pet{{
			Name:      "Pet1",
			CompanyID: &companyID,
		}},
	}

	err := DB.Create(&user).Error
	if err == nil {
		t.Errorf("should set the following error: Error 1452: Cannot add or update a child row: a foreign key constraint fails (`gorm`.`pets`, CONSTRAINT `fk_pets_company` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`))")
	}
}
