package main

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

const NR_BENCH_ITEMS = 10000

func prepareBenchTable() {

	var recordCount int64

	if err := DB.Table("bench_tables").Count(&recordCount).Error; err != nil {
		panic(err)
	}

	if recordCount == NR_BENCH_ITEMS {
		// fmt.Printf("bench db is already initialized, not reinitializing it\n")
		return
	}

	DB.Transaction(func(tx *gorm.DB) error {
		rand.Seed(1)

		if err := DB.Exec(`DELETE FROM bench_tables`).Error; err != nil {
			panic(err)
		}

		for i := 0; i < NR_BENCH_ITEMS; i++ {
			ent := BenchTable{
				BenchItem: i,
				Bool1:     rand.Int31n(10) < 7,
				Bool2:     rand.Int31n(10) < 7,
				Bool3:     rand.Int31n(10) < 7,
				Bool4:     rand.Int31n(10) < 7,
				String1:   fmt.Sprintf("%f", rand.Float64()),
				String2:   fmt.Sprintf("%f", rand.Float64()),
				String3:   fmt.Sprintf("%f", rand.Float64()),
				String4:   fmt.Sprintf("%f", rand.Float64()),
				Int1:      rand.Int(),
				Int2:      rand.Int(),
				Int3:      rand.Int(),
				Int4:      rand.Int(),
			}

			if err := DB.Create(&ent).Error; err != nil {
				panic(err)
			}
		}
		return nil
	})

}

func TestFetchAll(t *testing.T) {
	prepareBenchTable()

	assert := assert.New(t)

	data := []BenchTable{}
	assert.NoError(DB.Model(&BenchTable{}).Find(&data).Error)

	assert.Len(data, NR_BENCH_ITEMS)
}

func BenchmarkFirst(b *testing.B) {
	prepareBenchTable()
	DB.Logger = DB.Logger.LogMode(logger.Warn)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		id := rand.Int63n(NR_BENCH_ITEMS)
		tbl := BenchTable{}

		if err := DB.Model(&BenchTable{}).Where("bench_item = ?", id).First(&tbl).Error; err != nil {
			panic(err)
		}
	}
}

func BenchmarkFetchAll(b *testing.B) {
	prepareBenchTable()
	DB.Logger = DB.Logger.LogMode(logger.Warn)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tbl := []BenchTable{}

		if err := DB.Model(&BenchTable{}).Find(&tbl).Error; err != nil {
			panic(err)
		}
	}
}
