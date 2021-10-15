package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Record struct {
	gorm.Model
	TaskID       uint64 `gorm:"Index"`
	BatchId      uint32 `gorm:"Index"`
	DstObjectId  uint64 `gorm:"Index"`
	SrcObjectIds string `gorm:"type:blob"` //用,分割objectid
}

func openDB(debug, singularTable bool) (*gorm.DB, error) {
	var err error
	var db *gorm.DB

	logLevel := logger.Warn
	if debug {
		fmt.Print("openDB with debug mode")
		logLevel = logger.Info
	}

	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logLevel,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	sqlconnstr := "root:root@tcp(localhost:3306)/testDB2?charset=utf8mb4&parseTime=true&loc=Local"
	db, err = gorm.Open(mysql.Open(sqlconnstr), &gorm.Config{
		Logger: logger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: singularTable,
		},
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		AllowGlobalUpdate:      true,
	})
	return db, err
}

func main() {

	db, err := openDB(true, false)
	if err != nil {
		fmt.Print("OpenDB failed, err=", err)
		return
	}

	sqlDB, e := db.DB()
	if e != nil {
		fmt.Print("Get Sql DB failed, err=", e)
		err = e
		return
	}

	if err = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;").AutoMigrate(&Record{}); err != nil {
		fmt.Print("create gatherstore database error ", err)
		sqlDB.Close()
		return
	}

	db = db.Model(&Record{})

	var results []Record
	query := db.Where("task_id=1 AND dst_object_id=1448200665468768285").Select("dst_object_id, src_object_ids").Order("id DESC").Find(&results)
	if query.Error != nil {
		return
	}

	// do something

	record := Record{
		TaskID:       1,
		BatchId:      28,
		DstObjectId:  1448200665468768285,
		SrcObjectIds: "1448482124259659796,1448200665468768285",
	}
	if err := db.Create(&record).Error; err != nil {
		fmt.Print("create err,err=", err)
		return
	}
}
