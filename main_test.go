package main

import (
	"math/rand"
	"testing"
	"time"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)
	var sameUser User
	DB.Where("id =?", user.ID).First(&sameUser)
	println(user.UpdatedAt.Format("2006-01-02 15:04:05"))
	println(sameUser.UpdatedAt.Format("2006-01-02 15:04:05"))

	totalItemCreated := 100
	for i := 0; i < totalItemCreated; i++ {
		rand.Seed(time.Now().UnixNano())

		var job AdSet
		job.AdsetName = randSeq(10)
		job.AdsetId = uint64(i + 1)
		job.Region = "[]"
		job.Device = "[]"
		job.Extra = "[]"
		DB.Create(&job)
		var sameJob AdSet
		DB.Where("adset_id = ? ", i+1).First(&sameJob)

		formatStr := "2006-01-02 15:04:05"

		if job.UpdatedAt.Format(formatStr) != sameJob.UpdatedAt.Format(formatStr) {
			t.Errorf("Expected %v got %v", sameJob.UpdatedAt.Format(formatStr), job.UpdatedAt.Format(formatStr))
		}
	}

}
