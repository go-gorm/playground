package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {

	require.NoError(t, DB.Exec(`DROP VIEW IF EXISTS first_foo`).Error)
	require.NoError(t, DB.Exec(`DROP TABLE IF EXISTS foo`).Error)

	err := DB.AutoMigrate(&Foo1{})
	require.NoError(t, err)

	err = DB.Exec(`CREATE VIEW first_foo AS SELECT name FROM foo LIMIT 1`).Error
	assert.NoError(t, err)

	// Now migrate table by changing type of column
	err = DB.AutoMigrate(&Foo2{})
	assert.NoError(t, err)

}

type Foo1 struct {
	ID   string
	Name string
	Age  int
}

type Foo2 struct {
	ID   string
	Name string `gorm:"type:varchar(255)"`
	Age  int
}

func (Foo1) TableName() string {
	return "foo"
}

func (Foo2) TableName() string {
	return "foo"
}
