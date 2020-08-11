package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}
	DB.Create(&user)

	company := Company{
		ID:   int(user.ID),
		Name: user.Name,
	}
	DB.Create(&company)

	i := new(int64)
	if err := DB.Table("users").Joins("INNER JOIN companies on companies.name = users.name").Where("users.name = ?", user.Name).Count(i).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}