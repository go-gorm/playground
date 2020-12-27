package main

import (
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

var DB *gorm.DB
var ReadDB *gorm.DB

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

		ReadDB = initReadDB()
		if DB.Dialector.Name() == "sqlite" {
			ReadDB.Exec("PRAGMA foreign_keys = ON")
		}
		ReadDB.Logger = ReadDB.Logger.LogMode(logger.Info)
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
	var err error
	allModels := []interface{}{&User{}, &Account{}, &Pet{}, &Company{}, &Toy{}, &Language{}}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(allModels), func(i, j int) { allModels[i], allModels[j] = allModels[j], allModels[i] })

	DB.Migrator().DropTable("user_friends", "user_speaks")

	if err = DB.Migrator().DropTable(allModels...); err != nil {
		log.Printf("Failed to drop table, got error %v\n", err)
		os.Exit(1)
	}

	if err = DB.AutoMigrate(allModels...); err != nil {
		log.Printf("Failed to auto migrate, but got error %v\n", err)
		os.Exit(1)
	}

	for _, m := range allModels {
		if !DB.Migrator().HasTable(m) {
			log.Printf("Failed to create table for %#v\n", m)
			os.Exit(1)
		}
	}
}

func initReadDB() (db *gorm.DB) {
	var err error

	dbDSN := os.Getenv("GORM_DSN")
	dbDSN2 := os.Getenv("GORM_DSN_READ")
	switch os.Getenv("GORM_DIALECT") {
	case "mysql":
		if dbDSN == "" {
			dbDSN = "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
		}
		if dbDSN2 == "" {
			dbDSN2 = "gorm:gorm@tcp(localhost:9911)/gorm?charset=utf8&parseTime=True&loc=Local"
		}
		db, err = gorm.Open(mysql.Open(dbDSN2), &gorm.Config{})
		if err == nil {
			err = db.Use(dbresolver.Register(dbresolver.Config{
				Sources:  []gorm.Dialector{mysql.Open(dbDSN)},
				Replicas: []gorm.Dialector{mysql.Open(dbDSN2)},
			}, "users"))
		}
	case "postgres":
		if dbDSN == "" {
			dbDSN = "user=gorm password=gorm host=localhost dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
		}
		if dbDSN2 == "" {
			dbDSN2 = "user=gorm password=gorm host=localhost dbname=gorm port=9921 sslmode=disable TimeZone=Asia/Shanghai"
		}
		db, err = gorm.Open(postgres.Open(dbDSN2), &gorm.Config{})
		if err == nil {
			err = db.Use(dbresolver.Register(dbresolver.Config{
				Sources:  []gorm.Dialector{postgres.Open(dbDSN)},
				Replicas: []gorm.Dialector{postgres.Open(dbDSN2)},
			}, "users"))
		}

	case "sqlserver":
		// CREATE LOGIN gorm WITH PASSWORD = 'LoremIpsum86';
		// CREATE DATABASE gorm;
		// USE gorm;
		// CREATE USER gorm FROM LOGIN gorm;
		// sp_changedbowner 'gorm';
		if dbDSN == "" {
			dbDSN = "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
		}
		if dbDSN2 == "" {
			dbDSN2 = "sqlserver://gorm:LoremIpsum86@localhost:9931?database=gorm"
		}
		db, err = gorm.Open(sqlserver.Open(dbDSN2), &gorm.Config{})
		if err == nil {
			err = db.Use(dbresolver.Register(dbresolver.Config{
				Sources:  []gorm.Dialector{sqlserver.Open(dbDSN)},
				Replicas: []gorm.Dialector{sqlserver.Open(dbDSN2)},
			}, "users"))
		}

	default:
		db, err = gorm.Open(sqlite.Open(filepath.Join(os.TempDir(), "gorm.db1")), &gorm.Config{})
		if err == nil {
			err = db.Use(dbresolver.Register(dbresolver.Config{
				Sources:  []gorm.Dialector{sqlite.Open(filepath.Join(os.TempDir(), "gorm.db"))},
				Replicas: []gorm.Dialector{sqlite.Open(filepath.Join(os.TempDir(), "gorm.db1"))},
			}, "users"))
		}
	}
	if err != nil {
		log.Printf("failed to connect database, got error %v\n", err)
		os.Exit(1)
	}

	if err = db.Migrator().DropTable(&User{}); err != nil {
		log.Printf("Failed to drop table, got error %v\n", err)
		os.Exit(1)
	}

	if err = db.AutoMigrate(&User{}); err != nil {
		log.Printf("Failed to auto migrate, but got error %v\n", err)
		os.Exit(1)
	}

	return
}
