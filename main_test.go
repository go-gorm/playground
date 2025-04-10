package main

import (
	"gorm.io/gorm"
	"testing"
	"time"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	preTime := time.Date(2025, 4, 10, 0, 0, 0, 0, time.UTC)
	var preId uint = 10
	user := User{
		Model: gorm.Model{
			ID:        preId,
			CreatedAt: preTime,
		},
		Name: "jinzhu",
	}

	DB.Create(&user)

	user2 := User{
		Model: gorm.Model{
			ID:        12,
			CreatedAt: time.Date(2025, 4, 11, 0, 0, 0, 0, time.UTC),
		},
		Name: "jinzhu",
	}

	err := DB.Debug().Select("*").Omit("id").Omit("created_at").Omit("name").
		Where("name = ?", "jinzhu").
		Updates(user2).Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	getUser := User{}
	err = DB.First(&getUser, user2.ID).Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	if getUser.ID != preId {
		t.Errorf("id changed!!!")
	}
	if getUser.CreatedAt != preTime {
		t.Errorf("created_at changed!!!")
	}
	/*
		2025/04/10 23:00:53 testing sqlite3...
		=== RUN   TestGORM

		2025/04/10 23:00:53 /Users/cruvie/cruvie-space/code-hub/open-source/playground/main_test.go:24
		[1.637ms] [rows:1] INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`name`,`age`,`birthday`,`company_id`,`manager_id`,`active`,`id`) VALUES ("2025-04-10 00:00:00","2025-04-10 23:00:53.669",NULL,"jinzhu",0,NULL,NULL,NULL,false,10) RETURNING `id`

		2025/04/10 23:00:53 /Users/cruvie/cruvie-space/code-hub/open-source/playground/main_test.go:36
		[0.154ms] [rows:0] UPDATE `users` SET `id`=12,`created_at`="2025-04-11 00:00:00",`updated_at`="2025-04-10 23:00:53.671",`deleted_at`=NULL,`age`=0,`birthday`=NULL,`company_id`=NULL,`manager_id`=NULL,`active`=false WHERE name = "jinzhu" AND `users`.`deleted_at` IS NULL AND `id` = 12

		2025/04/10 23:00:53 /Users/cruvie/cruvie-space/code-hub/open-source/playground/main_test.go:41 record not found
		[0.043ms] [rows:0] SELECT * FROM `users` WHERE `users`.`id` = 12 AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1
		    main_test.go:43: Failed, got error: record not found
		    main_test.go:46: id changed!!!
		    main_test.go:49: created_at changed!!!
		--- FAIL: TestGORM (0.00s)

		FAIL

		进程 已完成，退出代码为 1
	*/
}
