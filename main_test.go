package main

import (
	"context"
	"errors"
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	ctx := context.Background()
	//	ctx, _ = context.WithTimeout(ctx, time.Nanosecond)

	ctx, cancelFunc := context.WithCancel(ctx)
	cancelFunc()

	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var result User
		if err := tx.First(&result, user.ID).Error; err != nil {
			t.Errorf("Failed, got error: %v", err)
		}
		user.Name = "foobar"
		user.ID = 0
		tx.Create([]User{user})

		return errors.New("some error")
	})

	t.Error(err)
}
