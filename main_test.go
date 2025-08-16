package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result User
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
