package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "test user"}

	r := DB.Create(&user)
	if r.Error != nil {
		t.Errorf("error: %v", r.Error)
	}

	pet := Pet{
		UserID: &user.ID,
		Name:   "Pet 1",
	}

	r = DB.Create(&pet)
	if r.Error != nil {
		t.Errorf("error: %v", r.Error)
	}

	pet = Pet{
		UserID: &user.ID,
		Name:   "Pet 2",
	}

	r = DB.Create(&pet)
	if r.Error != nil {
		t.Errorf("error: %v", r.Error)
	}

	r = DB.Where(Pet{
		UserID: &user.ID,
	}).Delete(Pet{})
	if r.Error != nil {
		t.Errorf("error: %v", r.Error)
	}

	if r.RowsAffected != 2 {
		t.Errorf("unexpected number of affected rows. expected: 2; got: %v", r.RowsAffected)
	}

	pet = Pet{
		UserID: &user.ID,
		Name:   "Pet 3",
	}

	r = DB.Create(&pet)
	if r.Error != nil {
		t.Errorf("error: %v", r.Error)
	}

	r = DB.Where(Pet{
		UserID: &user.ID,
	}).Delete(Pet{})
	if r.Error != nil {
		t.Errorf("error: %v", r.Error)
	}

	if r.RowsAffected != 1 {
		t.Errorf("unexpected number of affected rows. expected: 1; got: %v", r.RowsAffected)
	}
}
