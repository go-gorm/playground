package main

import (
	"database/sql"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "de1987"}
	DB.Create(&user)

	account := Account{Number: "123456", UserID: sql.NullInt64{Int64: int64(user.ID)}}
	DB.Create(&account)

	var recoverUser User
	if err := DB.First(&recoverUser, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	var recoverAccount Account
	if err := DB.First(&recoverAccount, account.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	DB.Where("id = ?", account.ID).Delete(&Account{})
	DB.Where("id = ?", user.ID).Delete(&User{})

	if err := DB.First(&recoverUser, user.ID).Error; err == nil {
		t.Errorf("User should be deleted")
	}

	if err := DB.First(&recoverAccount, account.ID).Error; err == nil {
		t.Errorf("Account should be deleted")
	}

	DB.Save(&account)

	if err := DB.First(&recoverAccount, account.ID).Error; err == nil {
		t.Errorf("Account should keep deleted")
	}
}
