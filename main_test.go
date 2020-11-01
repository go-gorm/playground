package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}
	pet := Pet{Name: "snow"}
	DB.Create(&user)
	DB.Create(&pet)
	
	var result User
	subQuery1 := DB.Model(&User{}).Select("name")
	subQuery2 := DB.Model(&Pet{}).Select("name")

	db.Table("(?) as u, (?) as p", subQuery1, subQuery2).Find(&User{})

	// SELECT * FROM (SELECT `name` FROM `users`) as u, (SELECT `name` FROM `pets`) as p
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
