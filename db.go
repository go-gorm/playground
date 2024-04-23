package main

import (
	"context"
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

var queries = make([]string, 0)
var customLogger *gormLogger

type gormLogger struct {
	def    logger.Interface
	record bool
}

func (l *gormLogger) LogMode(level logger.LogLevel) logger.Interface {
	l.def = l.def.LogMode(level)
	return l
}

func (l *gormLogger) Info(ctx context.Context, s string, i ...interface{}) {
	if l.record {
		queries = append(queries, s)
	}

	l.def.Info(ctx, s, i...)
}

func (l *gormLogger) Warn(ctx context.Context, s string, i ...interface{}) {
	if l.record {
		queries = append(queries, s)
	}

	l.def.Warn(ctx, s, i...)
}

func (l *gormLogger) Error(ctx context.Context, s string, i ...interface{}) {
	if l.record {
		queries = append(queries, s)
	}

	l.def.Error(ctx, s, i...)
}

func (l *gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	l.def.Trace(ctx, begin, func() (string, int64) {
		sql, rowsAffected := fc()
		if l.record {
			queries = append(queries, sql)
		}

		return sql, rowsAffected
	}, err)
}

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
	customLogger = &gormLogger{
		def: logger.Default,
	}

	cfg := gorm.Config{
		Logger: customLogger,
	}
	switch os.Getenv("GORM_DIALECT") {
	case "mysql":
		log.Println("testing mysql...")
		if dbDSN == "" {
			dbDSN = "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
		}
		db, err = gorm.Open(mysql.Open(dbDSN), &cfg)
	case "postgres":
		log.Println("testing postgres...")
		if dbDSN == "" {
			dbDSN = "user=gorm password=gorm host=localhost dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
		}
		db, err = gorm.Open(postgres.Open(dbDSN), &cfg)
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
		db, err = gorm.Open(sqlserver.Open(dbDSN), &cfg)
	default:
		log.Println("testing sqlite3...")
		db, err = gorm.Open(sqlite.Open(filepath.Join(os.TempDir(), "gorm.db")), &cfg)
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
	allModels := []interface{}{&User{}, &Company{}, Addresses{}}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(allModels), func(i, j int) { allModels[i], allModels[j] = allModels[j], allModels[i] })

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
