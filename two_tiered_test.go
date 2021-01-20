package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoTieredEntryDb(t *testing.T) {
	var ctted TwoTieredEntryDb

	e := &ctted.TwoTieredEntry
	e.EntryPK = 123
	e.Links = map[string]int64{"one": 1, "two": 2, "three": 3}

	result := DB.Create(&ctted)
	assert.NoError(t, result.Error)

	var ltted TwoTieredEntryDb
	result = DB.Preload("LinksDb").Find(&ltted, 123)
	assert.NoError(t, result.Error)
	assert.Equal(t, TwoTieredEntryDb{
		TwoTieredEntry: TwoTieredEntry{
			EntryPK: 123,
			Links:   map[string]int64{"one": 1, "two": 2, "three": 3},
		},
	}, ltted)
}
