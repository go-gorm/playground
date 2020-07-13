package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type User struct {
	ID        uint   `gorm:"type:mediumint;primary_key;AUTO_INCREMENT;unsigned" json:"ID"`
	Username  string `gorm:"type:varchar(30);not null" json:"username"`
	Password  string `gorm:"type:varchar(255)" json:"password"`
	Rank 	  string `gorm:"type:varchar(255)" json:"rank"`
	Email     string `gorm:"type:varchar(255)" json:"email"`
}

func (u *User) Get(db *gorm.DB) error {
	result := db.Debug().Where(&u).First(&u)
	if result.RecordNotFound() {
		return nil
	}
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func TestGORM(t *testing.T) {
	DB.AutoMigrate(&User{})
	user := User{
		Username: "test",
		Password: "test",
		Rank: "0",
		Email: "admin@localhost.com",
	}
	
	DB.Create(&user)
	
	find := User{
		Username: "test",
	}
	err := find.Get(DB)
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	if find.ID == 0 {
		t.Errorf("Failed, did not get correct struct %v", find)
	}
}
