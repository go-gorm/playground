package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: mysql

func TestGORM(t *testing.T) {
	songID := "003dc512121c4dfb898184565b92e8e8"

	artistsNameArr := make([]string, 0)

	artistModel := &Artist{}
	err := DB.Debug().Model(artistModel).
		Select([]string{"artists.name"}).
		Joins("left join song_artists on artists.id = song_artists.artist_id").
		Where("song_artists.song_id = ?", songID).
		Scan(artistsNameArr).Error
	if err != nil {
		panic(err)
	}
	// user := User{Name: "jinzhu"}

	// DB.Create(&user)

	// var result User
	// if err := DB.First(&result, user.ID).Error; err != nil {
	// 	t.Errorf("Failed, got error: %v", err)
	// }

}
