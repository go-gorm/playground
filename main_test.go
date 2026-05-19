package main

import (
	"testing"

	"gorm.io/playground/models"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestWithTX(t *testing.T) {
	users := []models.User{
		{Name: "jinzhu", Active: true},
		{Name: "kaaris", Active: false},
		{Name: "vald", Active: false},
	}

	DB.Create(users)

	tx := DB.Model(&models.User{})
	tx = tx.Where("age = 0")

	tx = tx.Where(
		tx.Or("Name = ?", users[0].Name).
			Or("Name = ?", users[1].Name),
	)
	// SELECT * FROM `users` WHERE (age = 0 OR Name = "jinzhu" OR Name = "kaaris" AND (age = 0 OR Name = "jinzhu" OR Name = "kaaris")) AND `users`.`deleted_at` IS NULL

	dest := []models.User{}
	results := tx.Find(&dest, nil)
	if results.Error != nil {
		t.Errorf("Failed, got error: %v", results)
	}

	if want, got := 2, len(dest); want != got {
		t.Errorf("Failed to filter, want %d, got %d", want, got)
	}
}

func TestWithDB(t *testing.T) {
	users := []models.User{
		{Name: "jinzhu", Active: true},
		{Name: "kaaris", Active: false},
		{Name: "vald", Active: false},
	}

	DB.Create(users)

	tx := DB.Model(&models.User{})
	tx = tx.Where("age = 0")

	tx = tx.Where(
		DB.Or("Name = ?", "jinzhu").
			Or("Name = ?", "kaaris"),
	)
	// SELECT * FROM `users` WHERE (Name = "jinzhu" OR Name = "kaaris") AND `users`.`deleted_at` IS NULL

	dest := []models.User{}
	results := tx.Find(&dest, nil)
	if results.Error != nil {
		t.Errorf("Failed, got error: %v", results)
	}

	if want, got := 2, len(dest); want != got {
		t.Errorf("Failed to filter, want %d, got %d", want, got)
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
