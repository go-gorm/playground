package main

import (
        "time"
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

        time.Sleep(5 * time.Second)

        r := DB.Exec("delete from users where created_at < now() - interval '? sec'", 2)
        if r.Error != nil {
           t.Errorf("Failed, got error: %v", r.Error)
        }
}
