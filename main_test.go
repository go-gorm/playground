package main

import (
	"gorm.io/gorm"
	"testing"

	"gorm.io/playground/models"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func WhereExpressionNotNil(t *testing.T) func(d *gorm.DB) *gorm.DB {
	return func(d *gorm.DB) *gorm.DB {
		session := d.Session(&gorm.Session{})
		// show 1st clause
		if session.Statement.Clauses["WHERE"].Expression == nil {
			t.Errorf("Should not be nil")
		}
		return d
	}
}

func TestGORM(t *testing.T) {
	user := models.User{Name: "jinzhu"}

	DB.Create(&user)

	var result models.User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	var userNil User
	query := DB.Scopes(WhereExpressionNotNil(t)).Where("Name NOT NULL")
	if query.Statement.Clauses["WHERE"].Expression == nil {
		t.Errorf("Should not be nil")
	}

	query.Find(&userNil)

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
