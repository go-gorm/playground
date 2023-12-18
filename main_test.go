package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	a := A{Name: "a"}
	result := DB.FirstOrCreate(&a)
	assert.NoError(t, result.Error)
	b := B{Name: "b"}
	result = DB.FirstOrCreate(&b)
	assert.NoError(t, result.Error)

	good := Good{A: a, B: b}
	x := DB.Create(&good)
	assert.NoError(t, x.Error)
	good2 := Good{A: a, B: b}
	// should get error because of primary key constraint
	x = DB.Create(&good2)
	assert.Error(t, x.Error)
	x = DB.FirstOrCreate(&good2, good2)
	assert.NoError(t, x.Error)

	bad := Bad{A: a, B: b}
	x = DB.Create(&bad)
	assert.NoError(t, x.Error)
	bad2 := Bad{A: a, B: b}
	// should get error because of primary key constraint, but doesn't, because primary key was not created
	x = DB.Create(&bad2)
	assert.Error(t, x.Error)
	x = DB.FirstOrCreate(&bad2, bad2)
	assert.NoError(t, x.Error)
}
