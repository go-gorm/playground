package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {

	var u User
	u.Name = "jinzhu"

	DB.Create(&u)

	err := DB.Model(&u).Update("Name", "Something Else").Error
	if err != nil {
		t.Errorf("Error updating record: %s", err)
	}

	/*var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}*/
}
