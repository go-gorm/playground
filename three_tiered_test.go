package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeTieredDb(t *testing.T) {
	var cttpd ThreeTieredParentDb
	cttpd.ParentPK = 1

	p := &cttpd.ThreeTieredParent
	p.Entries = []*ThreeTieredEntry{{
		EntryPK: 123,
		Links:   map[string]int64{"one": 1, "two": 2, "three": 3},
	}}

	result := DB.Create(&cttpd)
	assert.NoError(t, result.Error)

	var lttpd ThreeTieredParentDb
	result = DB.Preload("EntriesDb.LinksDb").Find(&lttpd, 1)
	assert.NoError(t, result.Error)
	assert.Equal(t, ThreeTieredParentDb{
		ThreeTieredParent: ThreeTieredParent{
			ParentPK: 1,
			Entries: []*ThreeTieredEntry{{
				EntryPK: 123,
				Links:   map[string]int64{"one": 1, "two": 2, "three": 3},
			}}},
	}, lttpd)
}
