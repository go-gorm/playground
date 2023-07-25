package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/clause"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	if DB.Dialector.Name() != "postgres" {
		t.Skip()
	}

	type WithDefaultValue struct {
		ID   int
		UUID string `gorm:"default:gen_random_uuid()"`
	}

	err := DB.Migrator().DropTable(&WithDefaultValue{})
	assert.Empty(t, err)
	err = DB.AutoMigrate(&WithDefaultValue{})
	assert.Empty(t, err)

	record := WithDefaultValue{ID: 1}
	err = DB.Clauses(clause.OnConflict{UpdateAll: true}).Create(&[]WithDefaultValue{record}).Error
	assert.Empty(t, err)

	assert.Equal(t, record.ID, 1)
	assert.NotEmpty(t, record.UUID)
}
