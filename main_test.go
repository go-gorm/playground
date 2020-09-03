package main

import (
	"encoding/json"
	"testing"

	"gorm.io/datatypes"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type Property struct {
	ID        int64       `gorm:"column:id;primaryKey"`
	PropKey   string      `gorm:"column:prop_key"`
	PropValue interface{} `gorm:"column:prop_value"`
}

type ValueDataA struct {
	VideoURL string `json:"video_url"`
}
type ValueDataB struct {
	ImageID int64 `json:"image_id"`
}

func TestGORM(t *testing.T) {
	DB.Migrator().AutoMigrate(&Property{})
	defer DB.Migrator().DropTable(&Property{})

	valuesA, _ := json.Marshal(&ValueDataA{VideoURL: "asdf.mp4"})
	testData := &Property{
		PropKey:   "video",
		PropValue: datatypes.JSON(valuesA),
	}

	if err := DB.Create(testData); err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
