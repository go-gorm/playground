package main

import (
	"context"
	"testing"
	"time"

	"gorm.io/playground/dal/query"
	"gorm.io/playground/model"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := model.User{Name: "jinzhu"}

	DB.Create(&user)

	var result model.User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	q := query.Use(DB)

	_, err := q.User.
		WithContext(context.TODO()).
		Where(q.User.Name.Eq(`jinzhu`)).
		UpdateSimple(q.User.UpdatedAt.Add(time.Second))

	if err != nil {
		t.Errorf(`Failed, got error: %v`, err)
	}

}
