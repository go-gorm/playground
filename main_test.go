package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TE-ST_DRIVERS: sqlite, mysql, postgres, sqlserver
// TEST_DRIVERS: postgres

func assertNoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	err := DB.Create(&user).Error
	assertNoError(t, err)

	company := Company{Name: "jinzhu corp"}

	err = DB.Create(&company).Error
	assertNoError(t, err)

	employment := Employment{
		Company: company,
		User:    user,
	}
	err = DB.Create(&employment).Error
	assertNoError(t, err)

	// delete should cascade and delete employment automatically
	err = DB.Delete(&User{}, user.ID).Error
	assertNoError(t, err)

	var count int64

	err = DB.Model(&User{}).Count(&count).Error
	assertNoError(t, err)
	if count != 0 {
		t.Errorf("Unexpected user count: %d", count)
	}

	err = DB.Model(&Employment{}).Count(&count).Error
	assertNoError(t, err)
	if count != 0 {
		t.Errorf("Unexpected employment count: %d", count)
	}
}
