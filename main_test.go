package main

import (
	"log"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("./user.db"), &gorm.Config{
		Logger: logger.New(log.Default(), logger.Config{
			LogLevel: logger.Info,
			Colorful: true,
		}),
	})

	if err != nil {
		t.Errorf("Failed to open db, got error: %v", err)
	}
	user := User{Username: "j2", PasswordHash: "test", Email: "t2"}

	err = db.AutoMigrate(&user)
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	// err = db.Save(&user).Error
	// if err != nil {
	// 	t.Errorf("Failed, got error: %v", err)
	// }
}
