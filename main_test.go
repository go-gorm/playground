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
// [0.578ms] [rows:1] INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`name`,`age`,`birthday`,`company_id`,`manager_id`,`active`) VALUES ("2021-12-23 14:26:08.079","2021-12-23 14:26:08.079",NULL,"jinzhu",0,NULL,NULL,NULL,false)
	query := DB.Model(User{})

	var c int64
	query.Where("name = 'something'").Count(&c)
// [0.060ms] [rows:1] SELECT count(*) FROM `users` WHERE name = 'something' AND `users`.`deleted_at` IS NULL

	var users []User
	query.Find(&users)
// [0.037ms] [rows:0] SELECT * FROM `users` WHERE name = 'something' AND `users`.`deleted_at` IS NULL

	if len(users) == 0 {
		t.Errorf("Found no one user")
	}
}
