package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	u := User{
		Name: "user",
		Token: Token{
			Content: "token",
		},
	}
	u1, err := saveUser(DB, &u)
	failOnError(t, err)

	expected := "token_encrypted"
	if u1.Token.Content != expected {
		t.Fatalf("expected %s, got %s", expected, u1.Token.Content)
	}

	u.Token.Content = "token2"
	u2, err := saveUser(DB, &u)
	failOnError(t, err)

	expected = "token2_encrypted"
	if u2.Token.Content != expected {
		t.Fatalf("expected %s, got %s", expected, u2.Token.Content)
	}
}

func saveUser(db *gorm.DB, u *User) (*User, error) {
	var newUser User
	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Session(&gorm.Session{FullSaveAssociations: true}).Save(u).Error; err != nil {
			return err
		}

		if err := tx.Preload("Token").First(&newUser, u.ID).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &newUser, nil
}

func failOnError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}
