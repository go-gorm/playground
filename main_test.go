package main

import (
	"testing"
	"uuid"

	"gorm.io/playground/models"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, postgres

func TestGORM(t *testing.T) {
	user := models.User{Name: "jinzhu"}

	DB.Create(&user)

	var result models.User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}

func TestUUIDOK(t *testing.T) {
	id, other := uuid.NewV4(), uuid.NewV4()

	t.Run("Create", func(t *testing.T) {
		if err := DB.Create(&models.Book{ID: id, Title: "foo"}).Error; err != nil {
			t.Fatalf("Failed, got error: %v", err)
		}
	})

	t.Run("Where with column", func(t *testing.T) {
		var book models.Book
		if err := DB.Where("id = ?", id).First(&book).Error; err != nil {
			t.Fatalf("got error: %v", err)
		}
		if book.ID != id {
			t.Errorf("expected %v, got %v", id, book.ID)
		}
	})

	t.Run("Where with struct", func(t *testing.T) {
		var book models.Book
		if err := DB.Where(&models.Book{ID: id}).First(&book).Error; err != nil {
			t.Fatalf("Failed, got error: %v", err)
		}
	})

	t.Run("Save", func(t *testing.T) {
		if err := DB.Save(&models.Book{ID: id, Title: "bar"}).Error; err != nil {
			t.Fatalf("Failed, got error: %v", err)
		}
	})

	t.Run("Find with []uuid.UUID", func(t *testing.T) {
		DB.Create(&models.Book{ID: other, Title: "baz"})

		var books []models.Book
		if err := DB.Find(&books, []uuid.UUID{id, other}).Error; err != nil {
			t.Fatalf("Failed, got error: %v", err)
		}
		if len(books) != 2 {
			t.Errorf("expected 2 rows, got %d", len(books))
		}
	})

	t.Run("Pluck", func(t *testing.T) {
		var ids []uuid.UUID
		if err := DB.Model(&models.Book{}).Pluck("id", &ids).Error; err != nil {
			t.Fatalf("Failed, got error: %v", err)
		}
		if len(ids) == 0 {
			t.Error("expected at least one id")
		}
	})
}

func TestUUIDKO(t *testing.T) {
	id := uuid.NewV4()
	if err := DB.Create(&models.Book{ID: id, Title: "foo"}).Error; err != nil {
		t.Fatalf("Failed, got error: %v", err)
	}

	t.Run("First", func(t *testing.T) {
		var book models.Book
		if err := DB.First(&book, id).Error; err != nil {
			t.Fatalf("Failed, got error: %v", err)
		}
	})

	t.Run("Find", func(t *testing.T) {
		var books []models.Book
		if err := DB.Find(&books, id).Error; err != nil {
			t.Fatalf("Failed, got error: %v", err)
		}
		if len(books) != 1 {
			t.Errorf("expected 1 book, got %d", len(books))
		}
	})

	t.Run("Delete", func(t *testing.T) {
		result := DB.Delete(&models.Book{}, id)
		if result.Error != nil {
			t.Fatalf("Failed, got error: %v", result.Error)
		}
		if result.RowsAffected != 1 {
			t.Errorf("expected RowsAffected=1, got %d", result.RowsAffected)
		}
	})
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
