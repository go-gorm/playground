package main

import (
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
	"log"
	"time"
)

type EventInt64 struct {
	Value     int32
	CreatedAt time.Time
}

// Both EventInt64 and EventString refer to one table "events"
func (e *EventInt64) TableName() string {
	return "events"
}

type EventString struct {
	Value     string
	CreatedAt time.Time
}

// Both EventInt64 and EventString refer to one table "events"
func (e *EventString) TableName() string {
	return "events"
}

func main() {
	dsn := "clickhouse://localhost:9915/default?dial_timeout=10s&read_timeout=20s&debug=1"
	DB, err := gorm.Open(clickhouse.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// create table "events" with String "Value"
	err = DB.
		Set("gorm:table_options", "ENGINE = MergeTree() PARTITION BY (toYYYYMM(created_at)) ORDER BY (created_at) SETTINGS index_granularity = 8192;").
		AutoMigrate(&EventString{})
	if err != nil {
		log.Fatal(err)
	}

	// insert record with String Value
	err = DB.Create(&EventString{
		Value:     "string value",
		CreatedAt: time.Now(),
	}).Error
	if err != nil {
		log.Fatal(err)
	}

	// migrate table "events" with Int64 "Value"
	err = DB.
		Set("gorm:table_options", "ENGINE = MergeTree() PARTITION BY (toYYYYMM(created_at)) ORDER BY (created_at) SETTINGS index_granularity = 8192;").
		AutoMigrate(&EventInt64{})
	if err != nil {
		log.Fatal(err)
	}
}
