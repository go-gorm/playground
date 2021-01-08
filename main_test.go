package main

import (
	"testing"
	"os"
	
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	//disable default transction due to mysql already done it inside.
	if os.Getenv("GORM_DIALECT") == "mysql" {
		db, _ := gorm.Open(mysql.Open("gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		})
		
		type Test struct {
			ID   int64 `gorm:"primaryKey"`
			Name string
		}

		if err = DB.AutoMigrate(&Test{}); err != nil {
			log.Printf("Failed to auto migrate, but got error %v\n", err)
			os.Exit(1)
		}
		
		user := Test{Name: "zzjin"}
		DB.Create(&user)
		
		if err := db.Where("id=?", user.ID).Delete(&Test{}).Error; err != nil {
			t.Errorf("Failed, got error: %v", err)
		}
		//again, panic
		if err := db.Where("id=?", user.ID).Delete(&Test{}).Error; err != nil {
			t.Errorf("Failed, got error: %v", err)
		}
	}
}
