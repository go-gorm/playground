package main

import (
	"testing"

	"gorm.io/gorm"
	"gorm.io/playground/models"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite

type UserExtended struct {
	models.User
	ExtraFoo int64
}

func TestGORM(t *testing.T) {
	user := models.User{Name: "jinzhu"}

	err := DB.Create(&user).Error
	if err != nil {
		t.Errorf("Create failed: %v", err)
	}

	var results []UserExtended
	err = DB.
		// setting model for magic table name
		Model(models.User{}).
		Select(`*, 42 AS extra_foo`).
		// taking results to different model with extra fields
		FindInBatches(&results, 100, func(tx *gorm.DB, batch int) error {
			// ...
			return nil
		}).
		Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
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
