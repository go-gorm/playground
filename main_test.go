package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver
type reportTable struct {
	gorm.Model
	Name string `gorm:"column:name" json:"name"`
}

func TestGORM(t *testing.T) {
	dsn := fmt.Sprintf("host=%s user=%s port=%d dbname=%s password=%s sslmode=disable",
		"127.0.0.1", "postgres", 5432, "test", "123456")
	pgDb, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}))
	if err != nil {
		t.Fatal(err)
	}
	pgDb = pgDb.Debug()
	pgDb.AutoMigrate(&reportTable{})
	test1 := &reportTable{Name: "test1,不插入id"}
	test2 := &reportTable{Name: "test2,插入id", Model: gorm.Model{ID: 2}}
	test3 := &reportTable{Name: "test3,不插入id"}
	if err := pgDb.Create(&test1).Error; err != nil {
		t.Fatal(err)
	}
	if err := pgDb.Create(&test2).Error; err != nil {
		t.Fatal(err)
	}
	if err := pgDb.Create(&test3).Error; err != nil {
		t.Fatal(err)
	}
}
