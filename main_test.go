package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	item := Item{
		Logo: "logo",
		Contents: []ItemContent{
			{
				LanguageCode: "en",
				Name:         "name",
			},
			{
				LanguageCode: "ar",
				Name:         "الاسم",
			},
		},
	}
	DB.Create(&item)

	DB.Model(&item).Association("Contents").Replace([]ItemContent{
		{
			LanguageCode: "en",
			Name:         "updated name",
		},
		{
			LanguageCode: "ar",
			Name:         "الاسم المحدث",
		},
		{
			LanguageCode: "fr",
			Name:         "le nom",
		},
	})

	var updatedItem Item
	DB.Preload("Contents").First(&updatedItem, item.ID)

	if len(updatedItem.Contents) != 3 {
		t.Errorf("expected 3 contents, got %d", len(updatedItem.Contents))
	}
}
