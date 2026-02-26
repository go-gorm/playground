package main

import (
	"gorm.io/gorm/schema"
	"log"
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

var (
	GormDB *gorm.DB
)

// Config gorm 自定义配置
func Config() *gorm.Config {
	return &gorm.Config{
		// 自定义日志记录器
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold: 200 * time.Millisecond, // 慢查询的门槛 200毫秒
			LogLevel:      logger.Info,            // 日志级别
			Colorful:      true,                   // 启用彩色输出
		}),
		// 命名策略
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",    // 表名前缀
			SingularTable: false, // true启用单数表名（false默认是复数表名）
		},
		// 设置为 true 后，GORM 在迁移时不会自动添加外键约束。适用于不需要外键约束的场景
		DisableForeignKeyConstraintWhenMigrating: true,
	}
}

func init() {
	pgsqlConfig := postgres.Config{
		DSN:                  "host=127.0.0.1 user=postgres password=admin dbname=go-gva port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		PreferSimpleProtocol: false,
	}
	if db, err := gorm.Open(postgres.New(pgsqlConfig), Config()); err != nil {
		return
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		GormDB = db
	}
}

//func init() {
//	var err error
//	if DB, err = OpenTestConnection(); err != nil {
//		log.Printf("failed to connect database, got error %v\n", err)
//		os.Exit(1)
//	} else {
//		sqlDB, err := DB.DB()
//		if err == nil {
//			err = sqlDB.Ping()
//		}
//
//		if err != nil {
//			log.Printf("failed to connect database, got error %v\n", err)
//		}
//
//		RunMigrations()
//		if DB.Dialector.Name() == "sqlite" {
//			DB.Exec("PRAGMA foreign_keys = ON")
//		}
//
//		DB.Logger = DB.Logger.LogMode(logger.Info)
//	}
//}

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

//func RunMigrations() {
//	var err error
//	allModels := []interface{}{&User{}, &Account{}, &Pet{}, &Company{}, &Toy{}, &Language{}}
//	rand.Seed(time.Now().UnixNano())
//	rand.Shuffle(len(allModels), func(i, j int) { allModels[i], allModels[j] = allModels[j], allModels[i] })
//
//	DB.Migrator().DropTable("user_friends", "user_speaks")
//
//	if err = DB.Migrator().DropTable(allModels...); err != nil {
//		log.Printf("Failed to drop table, got error %v\n", err)
//		os.Exit(1)
//	}
//
//	if err = DB.AutoMigrate(allModels...); err != nil {
//		log.Printf("Failed to auto migrate, but got error %v\n", err)
//		os.Exit(1)
//	}
//
//	for _, m := range allModels {
//		if !DB.Migrator().HasTable(m) {
//			log.Printf("Failed to create table for %#v\n", m)
//			os.Exit(1)
//		}
//	}
//}
