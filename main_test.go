package main

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user1 := User{Name: "jinzhu1", Age: 21}
	user2 := User{Name: "jinzhu2", Age: 22}
	user3 := User{Name: "jinzhu3", Age: 21}
	user4 := User{Name: "jinzhu4", Age: 21}

	DB.Create([]User{user1, user2, user3, user4})

	//Working as expected
	t.Run("Old API: Pluck", func(t *testing.T) {
		userIds := []uuid.UUID{}
		tx := DB.Model(&User{}).Where("age = ?", 21).Pluck("id", &userIds)
		if tx.Error != nil {
			t.Fatalf("Failed to pluck user IDs: %v", tx.Error)
		}

		if len(userIds) != 3 {
			t.Fatalf("Expected 3 users with age 21, got %d", len(userIds))
		}
	})

	//Working as expected
	t.Run("Old API: Select + Scan", func(t *testing.T) {
		userIds := []uuid.UUID{}
		tx := DB.Model(&User{}).Select("id").Where("age = ?", 21).Scan(&userIds)
		if tx.Error != nil {
			t.Fatalf("Failed to pluck user IDs: %v", tx.Error)
		}

		if len(userIds) != 3 {
			t.Fatalf("Expected 3 users with age 21, got %d", len(userIds))
		}
	})

	//Err: sql: Scan error on column index 0, name "id": unsupported Scan, storing driver.Value type string into type *[]uuid.UUID
	t.Run("New API: Select + Scan", func(t *testing.T) {
		userIds := []uuid.UUID{}
		ctx := context.Background()
		txErr := gorm.G[User](DB).Select("id").Where(User{Age: 21}).Scan(ctx, &userIds)
		if txErr != nil {
			t.Fatalf("Failed to pluck user IDs: %v", txErr)
		}

		if len(userIds) != 3 {
			t.Fatalf("Expected 3 users with age 21, got %d", len(userIds))
		}
	})
}
