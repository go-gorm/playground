package main

import (
	"fmt"
	"testing"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	err := DB.AutoMigrate(&DynamoDB{})
	if err != nil {
		panic("failed to auto-migrate table")
	}

	indexes := []string{
		"DROP INDEX IF EXISTS idx_unique;",
		"CREATE UNIQUE INDEX idx_unique ON dynamodb (hash_key, sort_key) WHERE NOT tombstone;",
	}
	for _, index := range indexes {
		err = DB.Exec(index).Error
		if err != nil {
			panic(fmt.Errorf("failed to apply index: %w", err))
		}
	}

	for i := 1; i <= 6; i++ {
		createMetadataVault(DB)
		fmt.Println("ran successfully for iteration: ", i)
	}
}

func createMetadataVault(db *gorm.DB) {
	var metadataModel = DynamoDB{
		HashKey: "hash",
		SortKey: "sort",
		Value:   "value",
	}
	result := db.
		Clauses(clause.OnConflict{
			Columns:     []clause.Column{{Name: "hash_key"}, {Name: "sort_key"}},
			TargetWhere: clause.Where{Exprs: []clause.Expression{clause.Eq{Column: "tombstone", Value: false}}},
			DoUpdates:   clause.AssignmentColumns([]string{"value"}),
		}).
		Create(&metadataModel)

	if result.Error != nil {
		panic("failed to create metadata vault")
	}
}
