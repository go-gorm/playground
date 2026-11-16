package main

import (
	"context"
	"testing"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type HookUser struct {
	ID         uint
	Name       string
	HookCalled bool `gorm:"-"`
}

func (u *HookUser) BeforeUpdate(tx *gorm.DB) error {
	u.HookCalled = true
	return nil
}

func TestGORM(t *testing.T) {
	ctx := context.Background()

	t.Run("traditional API invokes pointer hook", func(t *testing.T) {
		user := HookUser{}
		tx := DB.Session(&gorm.Session{
			DryRun:                 true,
			SkipDefaultTransaction: true,
		}).WithContext(ctx).Model(&user).
			Where("id = ?", 1).
			Updates(map[string]any{"name": "updated"})
		if tx.Error != nil {
			t.Fatal(tx.Error)
		}
		if !user.HookCalled {
			t.Fatal("BeforeUpdate was not called")
		}
	})

	t.Run("generics Set Update supports pointer hook", func(t *testing.T) {
		db := DB.Session(&gorm.Session{
			DryRun:                 true,
			SkipDefaultTransaction: true,
		})
		_, err := gorm.G[HookUser](db).
			Where("id = ?", 1).
			Set(clause.Assignments(map[string]any{"name": "updated"})).
			Update(ctx)
		if err != nil {
			t.Fatalf("expected pointer hook to work as in traditional API, got: %v", err)
		}
	})
}
