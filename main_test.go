package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	c1 := Company{Name: "Test"}
	res := DB.Create(&c1)
	require.Nil(t, res.Error)

	c2 := Company{}
	DB.Create(&c2)
	require.Nil(t, res.Error)

	var companies []Company
	res = DB.Model(Company{}).Find(&companies)
	require.Nil(t, res.Error)
	require.Equal(t, 2, len(companies))

	require.Equal(t, "Test", companies[0].Name)
	require.Equal(t, "", companies[1].Name)
}
