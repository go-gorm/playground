package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type NullableStringType struct {
	Value string `gorm:"default:null"`
}

func TestGORM(t *testing.T) {
	DB.AutoMigrate(&NullableStringType{})
	value := NullableStringType{Value: "jinzhu"}
	value2 := NullableStringType{}
	values := []NullableStringType{value, value2}

	tx := DB.Create(&values)
	if tx.Error != nil {
		// Fails for sqlite - near "DEFAULT": syntax error
		// [0.210ms] [rows:0] INSERT INTO `nullable_string_types` (`value`) VALUES ("jinzhu"),(DEFAULT) RETURNING `value`
		t.Fail()
	}
}
