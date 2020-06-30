package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {

	user := User{
		Avatar:   "1.png",
		Nickname: "test1",
	}
	DB.FirstOrCreate(&user)

	comments := []Comment{
		{
			UserId:  user.ID,
			Content: "comment.content-1",
		},
		{
			UserId:  user.ID,
			Content: "comment.content-2",
		},
	}
	DB.Create(comments)

	var rows []Comment

	// 查询所有前10条评论 关联相对应的用户
	if err := DB.Limit(10).Preload("User").Find(&rows).Error; err != nil {
		t.Error(err.Error())
	} else {
		for _, row := range rows {
			t.Log(row)
			t.Log(row.User)
		}
	}
}
