package main

import (
	"log"
	"os"
	"path/filepath"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func init() {
	var err error
	if DB, err = OpenTestConnection(); err != nil {
		log.Printf("failed to connect database, got error %v\n", err)
		os.Exit(1)
	} else {
		sqlDB, err := DB.DB()
		if err == nil {
			err = sqlDB.Ping()
		}

		if err != nil {
			log.Printf("failed to connect database, got error %v\n", err)
		}

		RunMigrations()
		if DB.Dialector.Name() == "sqlite" {
			DB.Exec("PRAGMA foreign_keys = ON")
		}

		DB.Logger = DB.Logger.LogMode(logger.Info)
	}
}

func OpenTestConnection() (db *gorm.DB, err error) {
	dbDSN := os.Getenv("GORM_DSN")
	switch os.Getenv("GORM_DIALECT") {
	case "mysql":
		log.Println("testing mysql...")
		if dbDSN == "" {
			dbDSN = "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
		}
		db, err = gorm.Open(mysql.Open(dbDSN), &gorm.Config{})
	case "postgres":
		log.Println("testing postgres...")
		if dbDSN == "" {
			dbDSN = "user=gorm password=gorm host=localhost dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
		}
		db, err = gorm.Open(postgres.Open(dbDSN), &gorm.Config{})
	case "sqlserver":
		// CREATE LOGIN gorm WITH PASSWORD = 'LoremIpsum86';
		// CREATE DATABASE gorm;
		// USE gorm;
		// CREATE USER gorm FROM LOGIN gorm;
		// sp_changedbowner 'gorm';
		log.Println("testing sqlserver...")
		if dbDSN == "" {
			dbDSN = "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
		}
		db, err = gorm.Open(sqlserver.Open(dbDSN), &gorm.Config{})
	default:
		log.Println("testing sqlite3...")
		db, err = gorm.Open(sqlite.Open(filepath.Join(os.TempDir(), "gorm.db")), &gorm.Config{})
	}

	if debug := os.Getenv("DEBUG"); debug == "true" {
		db.Logger = db.Logger.LogMode(logger.Info)
	} else if debug == "false" {
		db.Logger = db.Logger.LogMode(logger.Silent)
	}

	return
}

func RunMigrations() {
	if err := DB.Migrator().DropTable(&Address{}, &User{}); err != nil {
		log.Print("failed to drop tables", err)
		os.Exit(1)
	}
	// run Address2 migration
	if err := DB.Migrator().AutoMigrate(&Address2{}); err != nil {
		log.Print("failed to migrate address 2: ", err)
		os.Exit(1)
	}
	// Run User migration
	err := DB.Migrator().AutoMigrate(&User{})
	if err != nil {
		log.Print("failed to migrate user", err)
		os.Exit(1)
	}
	// check new address2 fields are there
	foundSecondField := false
	foundUserId := false
	userIdUnique := false
	columns, err := DB.Migrator().ColumnTypes(&Address2{})
	if err != nil {
		log.Print("could not get columns for address 2: ", err)
		os.Exit(1)
	}
	for _, column := range columns {
		if column.Name() == "user_id" {
			foundUserId = true
			isUnique, isOkay := column.Unique()
			if !isOkay {
				log.Print("found error checking is field is unique")
				os.Exit(1)
			}
			userIdUnique = isUnique
		} else if column.Name() == "second_field" {
			foundSecondField = true
		}
	}
	if !foundSecondField {
		log.Print("could not find second field---did address2 affect the same table?")
		os.Exit(1)
	}
	if !foundUserId {
		log.Print("could not find second field---did address2 affect the same table?")
		os.Exit(1)
	}
	// check column with userId does not have non-null conditional on it
	if userIdUnique {
		log.Print("found that user id was unique, which doesn't match what we saw before")
		os.Exit(1)
	}
	// drop address table
	// run just address2 migration
	// check that column with user id is unique

}
