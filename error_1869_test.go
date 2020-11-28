package main

import (
	"testing"
)


// User struct is a row record of the facebook_profile table in the digsty database
type CustomUser struct {
	ID uint64 `gorm:"primaryKey;column:id;type:bigint unsigned AUTO_INCREMENT;" json:"id"`
	Name string `gorm:"column:name;uniqueIndex;type:varchar(255) NOT NULL DEFAULT '';size:255;" json:"name"`
}

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	
	
	user := CustomUser{Name: "jinzhu"}

	DB.
	Clauses(clause.OnConflict{DoNothing: true }).
	Create(&user)
	
	
	user := CustomUser{Name: "jinzhu"}
	DB.
	Clauses(clause.OnConflict{DoNothing: true }).
	Create(&user)
	

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
