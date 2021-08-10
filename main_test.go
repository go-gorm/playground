package main

import (
	"testing"

	"github.com/google/uuid"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {

	id := uuid.New().String()
	DB.AutoMigrate(&BaseModelWithOneEmbedded{}, &BaseModelWithTwoEmbedded{})

	baseModelWithOneEmbedded := BaseModelWithOneEmbedded{
		BaseModelId: id,
		ChildOne:    ChildOne{},
	}

	DB.Create(&baseModelWithOneEmbedded)

	var result1 BaseModelWithOneEmbedded
	if err := DB.First(&result1, "BASE_MODEL_ID = ?", baseModelWithOneEmbedded.BaseModelId).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	if result1.ChildOne.Column1 != "WORKING" {
		t.Errorf("BeforeCreate not called on ChildOne for BaseModelWithOneEmbedded")
	} else {
		t.Log("Working for one embedded")
	}

	//BaseModel with two embedded model
	baseModelWithTwoEmbedded := BaseModelWithTwoEmbedded{
		BaseModelId: id,
		ChildOne:    ChildOne{},
		ChildTwo:    ChildTwo{},
	}

	DB.Create(&baseModelWithTwoEmbedded)

	var result2 BaseModelWithTwoEmbedded
	if err := DB.First(&result2, "BASE_MODEL_ID = ?", baseModelWithTwoEmbedded.BaseModelId).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	if result2.ChildOne.Column1 != "WORKING" {
		t.Errorf("BeforeCreate not called on ChildOne for BaseModelWithTwoEmbedded")
	}
	if result2.ChildTwo.Column2 != "WORKING" {
		t.Errorf("BeforeCreate not called on ChildTwo for BaseModelWithTwoEmbedded")
	}
}
