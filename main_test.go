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

	err := DB.Debug().Select("*").Omit("id").Omit("created_at").Omit("name").
		Where("name = ?", "jinzhu").
		Updates(user).Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	/*
	   [0.883ms] [rows:1] UPDATE `users` SET `id`=1,`created_at`="2025-04-10 22:50:54.487",`updated_at`="2025-04-10 22:50:54.488",`deleted_at`=NULL,`age`=0,`birthday`=NULL,`company_id`=NULL,`manager_id`=NULL,`active`=false WHERE name = "jinzhu" AND `users`.`deleted_at` IS NULL AND `id` = 1
	   --- PASS: TestGORM (0.00s)
	*/
}
