package main

import (
	"strconv"
	"testing"

	"gorm.io/gorm"
	"gorm.io/playground/models"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite

func TestFindInBatchesNonPositiveBatchSize(t *testing.T) {
	user := models.User{Name: "invalid_batch_size"}
	if err := DB.Create(&user).Error; err != nil {
		t.Fatal(err)
	}

	for _, batchSize := range []int{0, -1} {
		t.Run(strconv.Itoa(batchSize), func(t *testing.T) {
			var users []models.User
			calls := 0
			defer func() {
				if value := recover(); value != nil {
					t.Errorf("FindInBatches panicked after %d callback calls: %v", calls, value)
				}
			}()

			err := DB.Where("id = ?", user.ID).FindInBatches(&users, batchSize, func(*gorm.DB, int) error {
				calls++
				return nil
			}).Error
			if err == nil {
				t.Error("expected an error for a non-positive batch size")
			}
			if calls != 0 {
				t.Errorf("expected no callback calls, got %d", calls)
			}
		})
	}
}

// func TestGORMGen(t *testing.T) {
// 	user := models.User{Name: "jinzhu2"}
// 	ctx := context.Background()

// 	gorm.G[models.User](DB).Create(ctx, &user)

// 	if u, err := gorm.G[models.User](DB).Where(g.User.ID.Eq(user.ID)).First(ctx); err != nil {
// 		t.Errorf("Failed, got error: %v", err)
// 	} else if u.Name != user.Name {
// 		t.Errorf("Failed, got user name: %v", u.Name)
// 	}
// }
