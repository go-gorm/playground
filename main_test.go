package main

import (
	"testing"

	"github.com/google/uuid"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: postgres

func TestGORM(t *testing.T) {
	tomID := uuid.New().String()
	DB.Create(&User{
		ID:   tomID,
		Name: "tom",
	})

	bobID := uuid.New().String()
	DB.Create(&User{
		ID:   bobID,
		Name: "bob",
	})

	DB.Create(&UserAvatar{
		ID:     uuid.New().String(),
		UserID: tomID,
		Size:   1337,
	})

	DB.Create(&UserAvatar{
		ID:     uuid.New().String(),
		UserID: bobID,
		Size:   2674,
	})

	var usersRoomAvatars usersRoomAvatars
	d := DB.Table("users").
		Select(
			"users.*",
			"user_avatars.id as avatar_id",
			"user_avatars.user_id as avatar_user_id",
			"user_avatars.size as avatar_size").
		Joins("left join user_avatars on user_avatars.user_id = users.id").
		Group("users.id, user_avatars.id")

	d.Scan(&usersRoomAvatars)

	for _, ua := range usersRoomAvatars {
		t.Logf("avatar ID: %s", ua.Avatar.ID)
	}
}
