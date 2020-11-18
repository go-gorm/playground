package main

import (
	"testing"
)

CustomUser{
		ID uint64 `gorm:"primaryKey;autoIncrement;column:id;type:bigint unsigned;" json:"id"`
}


// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	DB.AutoMigrate(&CustomUser{})
}
