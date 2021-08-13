package main

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := &User{
		ID:   uuid.New().String(),
		Name: "tom",
	}

	DB.Create(user)

	now := time.Now()
	login := &Login{
		ID:       uuid.New().String(),
		Location: "New York",
		Time:     &now,
	}
	DB.Create(login)

	DB.Model(&User{}).Where("id = ?", user.ID).Update("last_login_id", login.ID)

	var result User
	if err := DB.Joins("LastLogin").Where("users.id = ?", user.ID).First(&result).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	t.Logf("employee name %s last login location %s", result.Name, result.LastLogin.Location)
}
