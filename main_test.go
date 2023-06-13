package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	err := DB.AutoMigrate(&Policy{})
	assert.Nil(t, err)

	// Create a new policy with all default settings
	policy1 := Policy{Name: "policy10"}
	result1 := DB.Create(&policy1)
	assert.Nil(t, result1.Error)
	assert.NotZero(t, policy1.ID)
	assert.False(t, policy1.BoolFalse)
	assert.True(t, policy1.BoolTrue)

	policy2 := Policy{Name: "policy20", BoolFalse: true, BoolTrue: false}
	result2 := DB.Create(&policy2)
	assert.Nil(t, result2.Error)
	assert.NotZero(t, policy2.ID)
	assert.True(t, policy2.BoolFalse)
	assert.False(t, policy2.BoolTrue) // NOT RIGHT HERE

	// Save
	policy3 := Policy{Name: "policy30"}
	result3 := DB.Save(&policy3)
	assert.Nil(t, result3.Error)
	assert.NotZero(t, policy3.ID)
	assert.False(t, policy3.BoolFalse)
	assert.True(t, policy3.BoolTrue)

	policy4 := Policy{Name: "policy40", BoolFalse: true, BoolTrue: false}
	result4 := DB.Save(&policy4)
	assert.Nil(t, result4.Error)
	assert.NotZero(t, policy4.ID)
	assert.True(t, policy4.BoolFalse)
	assert.False(t, policy4.BoolTrue) // NOT RIGHT HERE
}
