package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func main() {
	dsn := "xxx:xxxx@tcp(xxxx)/xxx?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("connect db get error: %v", err)
		return
	}
	srcDb, err := db.DB()
	srcDb.SetMaxIdleConns(1)
	srcDb.SetMaxOpenConns(1)
	srcDb.SetConnMaxLifetime(time.Hour)

	db.Connection(func(dbs *gorm.DB) error {
		if result := dbs.Exec("create temporary table temp_order_weixin as select * from `order` where bank = 'weixin' AND id > 1"); result.Error != nil {
			log.Fatalf("create temp table get error: %v", result.Error)
			return result.Error
		}

		var count int64
		if result := dbs.Table("temp_order_weixin as o").
			Select("count(o.id)").
			Where("o.id > ?", 1).
			Count(&count); result.Error != nil {
			log.Fatalf("count get error: %v", result.Error)
			return result.Error
		}
		return nil
	})
}
