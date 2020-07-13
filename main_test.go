package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: mysql

type User2 struct {
	ID        uint   `gorm:"type:mediumint;primary_key;AUTO_INCREMENT;unsigned" json:"ID"`
	Username  string `gorm:"type:varchar(30);not null" json:"username"`
	Password  string `gorm:"type:varchar(255)" json:"password"`
	Rank 	  string `gorm:"type:varchar(255)" json:"rank"`
	Email     string `gorm:"type:varchar(255)" json:"email"`
}



func TestGORM(t *testing.T) {
	DB.AutoMigrate(&User2{})
	user := User2{
		Username: "test",
		Password: "test",
		Rank: "0",
		Email: "admin@localhost.com",
	}
	
	DB.Create(&user)
	
	find := User2{
		Username: "test",
	}
	err := DB.Debug().Where(&find).First(&find)
	if err != nil {
		t.Errorf("Failed, got error: %s", err.Error())
		return
	}
	if find.ID == 0 {
		t.Errorf("Failed, did not get correct struct %v", find)
	}
}
