package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm/clause"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: postgres

func TestGORM(t *testing.T) {

	db := DB

	db = db.Omit(clause.Associations)

	users := make([]User, 0)
	err := db.Find(&users).Error
	require.NoError(t, err)

	assert.Equal(t, "", db.Statement.Table)
}
