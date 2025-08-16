package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu", Age: 30}

	DB.Create(&user)

	type result struct {
		Name      string
		IsExample bool
		Age       int
	}
	var results []result
	err := DB.Table("users").Select("name = 'test example.com' as is_example", "age").Scan(&results).Error
	require.NoError(t, err)
	assert.False(t, results[0].IsExample)
	assert.Equal(t, results[0].Age, 30)

	// This time, an at-sign character is used in the select string. This causes the "age" string to not be included
	// as a column in select, thus resulting in age being a default zero.
	err = DB.Table("users").Select("name = 'test@example.com' as is_example", "age").Scan(&results).Error
	assert.False(t, results[0].IsExample)
	assert.Equal(t, results[0].Age, 30)

	require.NoError(t, err)
}
