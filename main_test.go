package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS:  postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{
		Name: "jinzhu",
		Age:  2,
	}
	anotherUser := User{
		Name: "with-company",
		Age:  1,
		EmbeddedCompany: &Company{
			ID:   2,
			Name: "company",
		}}

	DB.Create(&user)
	DB.Create(&anotherUser)

	var result []User
	err := DB.
		Order("age DESC"). // make sure user with company is always loaded first
		Find(&result).
		Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	foundUser := result[0]

	if foundUser.Name != "jinzhu" {
		t.Errorf("jhinzu should be loaded first because of the order by age DESC clause")
	}

	if foundUser.EmbeddedCompany != nil && foundUser.EmbeddedCompany.Name != "" {
		t.Errorf("user jinzhu should not have embedded company: %+v", foundUser.EmbeddedCompany)
	}
}
