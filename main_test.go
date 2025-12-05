package main

import (
	"testing"

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

	sql1 := `select name, age, birthday, active from users where name ~ '(zh){0,}.*' and name like ?`
	if rows, err := DB.Raw(sql1, "%in%").Rows(); err != nil {
		t.Errorf("Failed, got error: %v", err)
	} else {
		t.Logf("Succeeded, got result: %v", rows.Next())
	}

	sql2 := `select name, age, birthday, active from users where name ~ '(zh)?.*' and name like ?`
	if rows, err := DB.Raw(sql2, "%in%").Rows(); err != nil {
		t.Errorf("Failed, got error: %v", err)
	} else {
		t.Logf("Succeeded, got result: %v", rows.Next())
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
