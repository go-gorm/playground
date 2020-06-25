package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu", Company: Company{
		Name: "xxx",
	}}

	DB.Create(&user)

	var result User
	if err := DB.Select("users.id, c.name as company_name").Joins("inner join companies as c on users.company_id =  c.id ").
		Table("users").First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	if result.CompanyName == "" {
		t.Errorf("The field CompanyName is not set")
	}
}
