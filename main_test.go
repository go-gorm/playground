package main

import (
	"context"
	"slices"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver
func TestGORM(t *testing.T) {
	arrayInput := []string{"A", "B", "C"}
	arraySave := []string{"A1", "B1", "C1"}

	data := Model{
		JsonField: JsonField{
			Array: slices.Clone(arrayInput),
		},
	}

	t.Run("Create", func(t *testing.T) {
		err := DB.WithContext(context.TODO()).Model(&Model{}).
			Create(&data).
			Error
		if err != nil {
			t.Errorf("Failed, create error: %v", err)
		}

		arrayCreate := data.JsonField.Array
		t.Logf("arrayCreate is %v", arrayCreate)
		if !slices.Equal(arrayCreate, arrayInput) {
			t.Errorf("Failed, json field arrayCreate %v not equal arrayInput %v", arrayCreate, arrayInput)
		}
	})

	t.Run("Find With SkipCustomMethod", func(t *testing.T) {
		var modelSkip *ModelSkip
		err := DB.WithContext(context.TODO()).Model(&Model{}).
			Where("id", data.ID).
			Find(&modelSkip).
			Error
		if err != nil {
			t.Errorf("Failed, get error: %v", err)
		}

		arraySkip := modelSkip.JsonField.Array
		t.Logf("arraySkip is %v", arraySkip)
		if !slices.Equal(arraySkip, arraySave) {
			t.Errorf("Failed, modelSkip field arraySkip %v not equal arraySave %v", arraySkip, arraySave)
		}
	})

	t.Run("Find", func(t *testing.T) {
		var row *Model
		err := DB.WithContext(context.TODO()).Model(&Model{}).
			Where("id", data.ID).
			Find(&row).
			Error
		if err != nil {
			t.Errorf("Failed, find error: %v", err)
		}

		arrayFind := row.JsonField.Array
		t.Logf("arrayFind is %v", arrayFind)
		if !slices.Equal(arrayFind, arrayInput) {
			t.Errorf("Failed, json field arrayFind %v not equal arrayInput %v", arrayFind, arrayInput)
		}
	})

	t.Run("Update", func(t *testing.T) {
		arrayUpdate := []string{"X", "Y", "Z"}
		arrayUpdateSave := []string{"X1", "Y1", "Z1"}

		err := DB.WithContext(context.TODO()).Model(&Model{}).
			Where("id", data.ID).
			Updates(Model{
				JsonField: JsonField{
					Array: slices.Clone(arrayUpdate),
				},
			}).
			Error
		if err != nil {
			t.Errorf("Failed, update error: %v", err)
		}

		t.Run("ModelSkip", func(t *testing.T) {
			var modelSkip *ModelSkip
			err = DB.WithContext(context.TODO()).Model(&Model{}).
				Where("id", data.ID).
				Find(&modelSkip).Error
			if err != nil {
				t.Errorf("Failed, get error: %v", err)
			}

			arraySkip := modelSkip.JsonField.Array
			t.Logf("arraySkip is %v", arraySkip)
			if !slices.Equal(arraySkip, arrayUpdateSave) {
				t.Errorf("Failed, json field %v arrayInput not equal arrayUpdateSave %v", arraySkip, arrayUpdateSave)
			}
		})
	})

}
