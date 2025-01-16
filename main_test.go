package main

import (
	"log"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite

type Order struct {
	tableName string
	Id        uint64 `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"`
	UserId    uint64 `gorm:"column:user_id;type:bigint unsigned;not null;uniqueIndex:unique_user,priority:1" json:"user_id"`
}

func (o *Order) TableName() string {
	return o.tableName
}

func TestGORM(t *testing.T) {
	if err := DB.Table("order-1").AutoMigrate(&Order{}); err != nil {
		t.Error(err)
	}

	if err := DB.Table("order-2").AutoMigrate(&Order{}); err != nil {
		t.Error(err)
	}
	values := []map[string]any{}
	DB.Table("sqlite_master").Where("type = ?", "index").Select("type, name, tbl_name").Find(&values)

	log.Println(values)
}
