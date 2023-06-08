package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: mysql

func TestGORM(t *testing.T) {
	// user := User{Name: "jinzhu"}
	// 使用many2many创建关联表后,又再次AutoMigrate创建关联表,因为之前创建了联合索引(primary key (user_id, friend_id)),导致关系表无法增加新的逐渐索引.
	err := DB.AutoMigrate(&User{}, &UserFriend{})
	failOnError(t, err)
}

type UserFriend struct {
	ID       uint64 `gorm:"primary_key"`
	UserID   uint64 `gorm:"not null"`
	FriendID uint64 `gorm:"not null"`
}

func failOnError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}
