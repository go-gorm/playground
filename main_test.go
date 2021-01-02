package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	require := require.New(t)
	user := User{Name: "jinhzu"}

	// require.NoError(DB.Migrator().DropTable(&RUser{}))
	// require.NoError(DB.Migrator().CreateTable(&RUser{}))

	require.NoError(DB.Create(&user).Error)
	var result User
	require.NoError(DB.First(&result, user.ID).Error)
}
