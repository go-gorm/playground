package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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
}

func TestTimeZone(t *testing.T) {
	birthday := time.Date(2023, 11, 12, 9, 7, 18, 0, time.UTC)
	user := User{Name: "TimeZone UTC", Birthday: &birthday}

	if err := DB.Create(&user).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
		t.FailNow()
	}

	birthdayResult := *user.Birthday
	assert.Equal(t, birthday, birthdayResult)

	userResult := User{}

	if err := DB.Where("name = ?", "TimeZone UTC").First(&userResult).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
		t.FailNow()
	}

	birthdayResult = *userResult.Birthday
	assert.Equal(t, birthday, birthdayResult)

}
