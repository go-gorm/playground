package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	type TestItem struct {
		gorm.Model
		Data string
	}

	err := DB.AutoMigrate(TestItem{})
	assert.Nil(t, err)

	err = DB.Unscoped().Delete(TestItem{}).Error
	assert.Nil(t, err)

	size := 70000 // Max is 65536 placeholders
	items := make([]TestItem, 0, size)
	for idx := 0; idx < size; idx++ {
		items = append(items, TestItem{
			Model: gorm.Model{ID: uint(idx + 1)}, Data: fmt.Sprintf("Data_%d", idx),
		})
	}

	err = DB.CreateInBatches(items, 13000).Error
	assert.Nil(t, err)

	// Throws error, maybe good idea to have DeleteInBatch(value interface{}, batchSize int) similar to CreateInBatch
	err = DB.Unscoped().Delete(items).Error
	assert.Nil(t, err)
}
