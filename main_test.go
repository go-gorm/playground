package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type Table struct{ Value int }

func TestGORM(t *testing.T) {
	ctx := context.Background()
	tx1 := DB.Begin()
	defer tx1.Rollback()

	require.NoError(t, tx1.AutoMigrate(&Table{}))
	require.NoError(t, tx1.Create(&Table{Value: 1}).Error)

	t.Run("query once", func(t *testing.T) {
		require.NoError(t, query(ctx, tx1, 1))
	})

	t.Run("query 5 times", func(t *testing.T) {
		require.NoError(t, query(ctx, tx1, 5))
	})
}

func query(ctx context.Context, tx *gorm.DB, times int) error {
	eg, ctx := errgroup.WithContext(ctx)

	for i := 0; i < times; i++ {
		eg.Go(func() error {
			return tx.WithContext(ctx).Transaction(func(tx2 *gorm.DB) error {
				return tx2.Take(&Table{}).Error
			})
		})
	}

	return eg.Wait()
}
