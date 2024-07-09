package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type Foo struct {
	ID     int
	BarID  string
	FooBar *Bar `gorm:"foreignKey:BarID;references:ID"`
}

type Bar struct {
	ID   string
	Type string
}

func TestGORM(t *testing.T) {
	err := DB.Migrator().DropTable(&Foo{}, &Bar{})
	require.NoError(t, err)

	err = DB.AutoMigrate(&Foo{}, &Bar{})
	require.NoError(t, err)

	bar := []*Bar{
		{ID: "Bar-first", Type: "first"},
		{ID: "Bar-second", Type: "second"},
	}
	require.NoError(t, DB.Create(bar).Error)

	foo := []Foo{
		{BarID: bar[0].ID, ID: 1},
		{BarID: bar[1].ID, ID: 2},
	}
	require.NoError(t, DB.Create(foo).Error)

	var out Foo

	// Select that does not populate end struct, but shows broken backward compatibility.
	builder := DB.Select("foos.*, (?)", DB.Table("bars").Select("type"))
	// builer := DB.Table("foos") would work.

	builder = builder.
		Preload("FooBar").
		Joins("FooBar")
	builder = builder.Model(Foo{})
	builder = builder.Where(`"FooBar".type IN ?`, []string{"first"})

	res := builder.Find(&out)
	require.NoError(t, res.Error)
	require.EqualValues(t, bar[0], out.FooBar)
}
