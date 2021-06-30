package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

var (
	mockName string
	globalT  *testing.T
)

func (data *User) BeforeSave(tx *gorm.DB) error {
	assert.Equal(globalT, mockName, data.Name)
	return nil
}

func TestGORM(t *testing.T) {
	globalT = t
	mockName = "jinzhu"
	user := User{Name: mockName}

	// first create a record
	DB.Create(&user)

	// update `name` field
	originName := mockName
	mockName = "zhanghuihuang"
	DB.Model(&User{}).Where("name = ?", originName).Updates(User{
		Name: mockName,
	})
	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
