package main

import (
	"log"
	"os"
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Session(&gorm.Session{}).Create(&user)

	var result User

	var DB2 *gorm.DB
	var err error
	if DB2, err = OpenTestConnection(); err != nil {
		log.Printf("failed to connect database, got error %v\n", err)
		os.Exit(1)
	}

	if err := DB2.First(&result, user.ID).Update("name", "jinzhu 2").Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
