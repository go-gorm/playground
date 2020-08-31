package main

import (
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/playground/models"
	v1 "gorm.io/playground/models/v1"
	v2 "gorm.io/playground/models/v2"
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
			dbDSN = "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
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
	var err error
	allModels := []interface{}{&models.User{}, &models.Account{}, &models.Pet{}, &models.Company{}, &models.Toy{}, &models.Language{}}
	rand.Seed(time.Now().UnixNano())

	DB.Migrator().DropTable("user_friends", "user_speaks")

	if err = DB.Migrator().DropTable(allModels...); err != nil {
		log.Printf("Failed to drop table, got error %v\n", err)
		os.Exit(1)
	}

	if err = v1.Migrate(DB.Migrator()); err != nil {
		os.Exit(1)
	}
	if err = v2.Migrate(DB.Migrator()); err != nil {
		os.Exit(1)
	}
}
