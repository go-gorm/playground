package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	samples := []*Sample{
		{
			SampleId:      "s-1",
			OtherSampleId: "ss-1",
		},
		{
			SampleId:      "s-2",
			OtherSampleId: "ss-2",
		},
		{
			SampleId:      "s-3",
			OtherSampleId: "ss-3",
		},
	}
	DB.AutoMigrate(&Sample{})
	DB.Create(&samples)

	result := []*Sample{}
	var ids []string
	if err := DB.Find(&result).Pluck("other_sample_id", &ids).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
