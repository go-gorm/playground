package main

import (
	"log"
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

		if err := DB.AutoMigrate(&Test{}); err != nil {
			log.Printf("Failed to auto migrate, got error %v\n", err)
			os.Exit(1)
		}
		
		tx := pkg.MysqlClient.Begin()
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()
		if err := tx.Error; err != nil {
			log.Printf("Failed to start transaction, got error %v\n", err)
			os.Exit(1)
		}
		
		if err := tx.Where("name=?", "zzjin").Delete(&Test{}).Error; err != nil {
			tx.Rollback()
			
			log.Printf("Failed to run one transaction, got error %v\n", err)
			os.Exit(1)
		}
		
		if err := tx.Create(&Test{Name: "zzjin"}).Error; err != nil {
			tx.Rollback()
			
			log.Printf("Failed to run one transaction, got error %v\n", err)
			os.Exit(1)
		}
		
		if err := tx.Commit().Error; err != nil {
			log.Printf("Failed to commit transaction, got error %v\n", err)
			os.Exit(1)
		}
		
		//again, panic
		if err := db.Where("name=?", "zzjin").Delete(&Test{}).Error; err != nil {
			t.Errorf("Failed, got error: %v", err)
		}
	}
}
