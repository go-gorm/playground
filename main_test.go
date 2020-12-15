package main

import (
	"gorm.io/gorm"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	subQuery := DB.Table("users").Where("name=?", user.Name).Select("id")

	err := DB.Model(&Pet{}).Create([]map[string]interface{}{
		{
			"name": "cat",
			"user_id": gorm.Expr("(?)", DB.Table("(?) as tmp", subQuery).Select("@uid:=id")),
		},
		{
			"name": "dog",
			"user_id": gorm.Expr("@uid"),
		},
	}).Error

	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
