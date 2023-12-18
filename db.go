package main

import (
	"fmt"
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
	SetSessionAttribute()
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

const setTenantIdQuery = "SET myapp.current_tenant_id = '%s'"
const getTenantIdQuery = "SELECT current_setting('myapp.current_tenant_id')"

func SetSessionAttribute() {
	var currentTenantId string
	fmt.Printf("FATA-010 DB Ptr is [%p]\n", DB)
	fmt.Printf("FATA-020 & DB Ptr is [%p]\n", &DB)

	from1Session1 := getSession(DB, "TENANT1")
	fmt.Printf("FATA-030 DB session 1 is [%p]\n", from1Session1)
	fmt.Printf("FATA-040 &DB session 1 is [%p]\n", &from1Session1)

	DB.Raw(getTenantIdQuery).Scan(&currentTenantId)
	fmt.Printf("FATA-050 dbPtr1 tenantId is [%s]\n", currentTenantId)

	currentTenantId = ""
	from1Session1.Raw(getTenantIdQuery).Scan(&currentTenantId)
	fmt.Printf("FATA-060 from1Session1 tenantId is [%s]\n", currentTenantId)

	from1Session2 := getSession(DB, "TENANT2")
	fmt.Printf("FATA-070 DB session 2 is [%p]\n", from1Session2)
	fmt.Printf("FATA-080 &DB session 2 is [%p]\n", &from1Session2)

	currentTenantId = ""
	from1Session1.Raw(getTenantIdQuery).Scan(&currentTenantId)
	fmt.Printf("FATA-120 from1Session1 tenantId is [%s]\n", currentTenantId)

	currentTenantId = ""
	from1Session2.Raw(getTenantIdQuery).Scan(&currentTenantId)
	fmt.Printf("FATA-130 from1Session2 tenantId is [%s]\n", currentTenantId)

	return
}

func getSession(dbPtr *gorm.DB, tenantId string) *gorm.DB {
	sessionPtr := dbPtr.Exec(fmt.Sprintf(setTenantIdQuery, tenantId)).Session(&gorm.Session{})
	return sessionPtr

}
