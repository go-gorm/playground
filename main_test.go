package main

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: postgres

func TestGORM(t *testing.T) {

	entity := &Organisation{
		Name: "Admin Organisation",
	}
	db := DB

	db = db.Begin()
	if db.Error != nil {
		panic("Failed to begin test TX " + db.Error.Error())
	}

	ShowSearchPath(t, db)

	users := make([]User, 0)
	err := db.Find(&users).Error
	require.NoError(t, err)

	assert.Equal(t, "", db.Statement.Table)
	err = db.Omit(clause.Associations).Create(entity).Error
	require.NoError(t, err)
}

func ShowSearchPath(t *testing.T, db *gorm.DB) {
	type SearchPath struct {
		SearchPath string `store:"search_path"`
	}
	sp := &SearchPath{}
	err := db.Debug().Raw("SHOW search_path").Scan(&sp).Error
	if err != nil {
		require.NoError(t, err)
	}
	db.Logger.Info(context.Background(), "search_path: %s", sp.SearchPath)
}
