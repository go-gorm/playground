package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
func TestMigrationPostgresXXXSerial(t *testing.T) {
	type TableStruct struct {
		Version int64 `gorm:"type:bigserial"`
	}

	if err := DB.AutoMigrate(&TableStruct{}); err != nil {
		assert.Fail(t, "Auto migrate failed")
	}

	if err := DB.AutoMigrate(&TableStruct{}); err != nil {
		assert.Fail(t, "Auto migrate failed")
	}
}
