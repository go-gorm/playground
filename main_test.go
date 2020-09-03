package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type Property struct {
	ID        int64       `gorm:"column:id;primaryKey"`
	PropKey   string      `gorm:"column:prop_key"`
	PropValue interface{} `gorm:"column:prop_value"`
}

type PropValueDataA struct {
	VideoURL string `json:"video_url"`
}
type PropValueDataB struct {
	ImageID int64 `json:"image_id"`
}

func TestGORM(t *testing.T) {
	DB.Migrator().AutoMigrate(&Property{})
	defer DB.Migrator().DropTable(&Property{})

	if err := DB.Create(&Property{}); err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
