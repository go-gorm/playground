package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite

func TestGORM(t *testing.T) {
	names := make([]string, 0)

	// Works perfectly, no soucy !
	if err := DB.Table("pets").Pluck("name", &names).Error; err != nil {
		t.Errorf("error: %s", err.Error())
	}

	// FIXME: Doesn't works as I expect
	tx := DB.Table("pets AS p").Pluck("p.name", &names)
	if err := tx.Error; err != nil {
		t.Errorf("error: %s\n tx generate 'SELECT p.name FROM `pet AS p`' instead of 'SELECT p.name FROM `pet` AS `p`'", err.Error())
	}

	// FIXME: Doesn't works as I expect
	tx = DB.Table("public.pets AS p").Pluck("p.name", &names)
	if err := tx.Error; err != nil {
		t.Errorf("error: %s\n tx generate 'SELECT p.name FROM `public.pet AS p`' instead of 'SELECT p.name FROM `public.pet` AS `p`'", err.Error())
	}
}
