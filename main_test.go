package main

import (
	"database/sql"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestInWithParenthesis(t *testing.T) {
	user := User{Name: "jinzhu"}
	if err := DB.Create(&user).Error; err != nil {
		// Does not fail
		t.Errorf("Failed, got error: %v", err)
	}

	acount := Account{Number: "foo", UserID: sql.NullInt64{
		Int64: int64(user.ID),
		Valid: true,
	}}
	DB.Create(&acount)

	var accounts []Account
	q := DB.Select("accounts.*").Table("accounts").Joins("left join users on accounts.user_id = users.id and accounts.id IN (?)", []uint{1, 2})
	q = q.Find(&accounts)

	if err := q.Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}

func TestInWithoutParenthesis(t *testing.T) {
	user := User{Name: "jinzhu"}
	if err := DB.Create(&user).Error; err != nil {
		// Does not fail
		t.Errorf("Failed, got error: %v", err)
	}

	acount := Account{Number: "foo", UserID: sql.NullInt64{
		Int64: int64(user.ID),
		Valid: true,
	}}
	DB.Create(&acount)

	var accounts []Account
	q := DB.Select("accounts.*").Table("accounts").Joins("left join users on accounts.user_id = users.id and accounts.id IN ?", []uint{1, 2})
	q = q.Find(&accounts)

	if err := q.Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}

func TestWhereInWithParenthesis(t *testing.T) {
	user := User{Name: "jinzhu"}
	if err := DB.Create(&user).Error; err != nil {
		// Does not fail
		t.Errorf("Failed, got error: %v", err)
	}

	acount := Account{Number: "foo", UserID: sql.NullInt64{
		Int64: int64(user.ID),
		Valid: true,
	}}
	DB.Create(&acount)

	var accounts []Account
	q := DB.Select("accounts.*").Table("accounts").Where("accounts.id IN (?)", []uint{1, 2})
	q = q.Find(&accounts)

	if err := q.Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
