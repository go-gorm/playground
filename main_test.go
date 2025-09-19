package main

import (
	"testing"

	"gorm.io/gorm"
	"gorm.io/playground/models"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := models.User{Name: "jinzhu"}

	DB.Create(&user)

	var result models.User
	if err := DB.First(&result, user.ID).Error; err != nil {
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

func TestGORMScopesSessionTransactionLeak(t *testing.T) {
	sqldb, err := DB.DB()
	if err != nil {
		t.Error(err)
	}

	before := sqldb.Stats().InUse

	DB.Scopes(func(s *gorm.DB) *gorm.DB {
		return s.Session(&gorm.Session{
			DryRun: true,
			// Initialized: true,
		})
	}).Where("name = 'jinzhu'").Delete(&models.User{})

	after := sqldb.Stats().InUse

	if before != after {
		t.Errorf("Failed, InUse connection %d != %d", before, after)
	}
}
