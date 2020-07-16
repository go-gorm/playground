package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{
		Name: "jinzhu",
		Languages: []Language{{
			Code: "US",
		}, {
			Code: "ES",
		}}}

	DB.Create(&user)

	var result []User
	if err := DB.Distinct("u.id, u.*").Table("user_speaks as s").Joins("inner join users as u on u.id = s.user_id").
		Where(" s.language_code ='US' or s.language_code ='ES'").
		Find(&result).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if len(result) > 1 {
		t.Errorf("Failed, Distinct clause does not work")
	}
}
