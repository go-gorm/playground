package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"path"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver
type User struct {
	gorm.Model
	Name      string
	Age       uint
	Active    bool
}

func TestGORM(t *testing.T) {
	dbPath := path.Join("database.db")

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
	})
	if err != nil {
		log.Println(err)
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal(err)
	}

	user := User{Name: "jinzhu"}

	db.Create(&user)


	var result User
	if err := db.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
