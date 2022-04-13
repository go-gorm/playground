package main

import (
	"sync"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

var (
	once = &sync.Once{}
)

const USERS = 100
const SINGLE_BATCH = 5
const LIMIT = 50

func generateDatas(count int) (users []User) {
	for i := 0; i < count; i++ {
		users = append(users, User{Name: faker.Username()})
	}

	return
}

func TestGORM(t *testing.T) {
	once.Do(func() {
		users := generateDatas(USERS)
		DB.Create(&users)
	})
	var result []User
	DB.Find(&result)

	require.Equal(t, USERS, len(result), "fetched users are less than created")
}

func TestLimitClausedFindInBatches(t *testing.T) {
	once.Do(func() {
		users := generateDatas(USERS)
		DB.Create(&users)
	})

	var records, singleBatch []User
	ret := DB.Limit(LIMIT).FindInBatches(&singleBatch, SINGLE_BATCH, func(tx *gorm.DB, batch int) error {
		records = append(records, singleBatch...)
		return nil
	})

	affectedRows := int(ret.RowsAffected)

	require.Equal(t, affectedRows, len(records), "the processed rows count should equal")

	// FIXME: BELOW SHOULD NOT PASS!
	// Below shows it will overwrite original limit clause
	require.NotEqual(t, LIMIT, affectedRows, "the size with the affected rows should not equal")
	require.Equal(t, USERS, affectedRows, "the size with the total users should equal")
}

func TestOffsetClausedFindInBatches(t *testing.T) {
	once.Do(func() {
		users := generateDatas(USERS)
		DB.Create(&users)
	})

	for _, offset := range []int{2, 3, 5, 7} {
		// it will make a below query
		// || offset | single || offset | single || ... |
		// ended up with a partial or full query in below forms:
		// - || ... || offs||et
		// - || ... || offset ||
		// - || ... || offset | sing||le
		// - || ... || offset | single ||
		realBatches := int(USERS / (offset + SINGLE_BATCH))
		totalRetrieved := realBatches * SINGLE_BATCH
		remainEntries := USERS - (offset+SINGLE_BATCH)*realBatches
		if remainEntries > offset {
			totalRetrieved += remainEntries - offset
		}

		var records, singleBatch []User
		ret := DB.Offset(offset).FindInBatches(&singleBatch, SINGLE_BATCH, func(tx *gorm.DB, batch int) error {
			records = append(records, singleBatch...)
			return nil
		})

		affectedRows := int(ret.RowsAffected)

		require.Equal(t, affectedRows, len(records), "the processed rows count should equal")

		// FIXME: BELOW SHOULD NOT PASS!
		// It will not retrieve all rows
		require.NotEqual(t, USERS, affectedRows, "the size with the affected rows should not equal")
		// It will retrieve only totalRetrieved rows
		require.Equal(t, totalRetrieved, affectedRows, "the size with the total users should equal")
	}
}
