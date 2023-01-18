package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	DB.Model(User{}).Create(map[string]interface{}{"name": "test"})
	DB.Model(User{}).Where("name = ?", "test").Update("name", "testa")

}
