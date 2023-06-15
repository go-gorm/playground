package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	jinzhu1 := User{Name: "jinzhu1", Age: 2}
	jinzhu2 := User{Name: "jinzhu2", Age: 2}
	userA := User{Name: "A", ManagerID: &jinzhu1.ID}
	userB := User{Name: "B", ManagerID: &jinzhu1.ID}
	userC := User{Name: "C", ManagerID: &jinzhu1.ID}
	userD := User{Name: "D", ManagerID: &jinzhu1.ID}

	DB.Create(&jinzhu1)
	DB.Create(&jinzhu2)
	DB.Create(&userA)
	DB.Create(&userB)
	DB.Create(&userC)
	DB.Create(&userD)

	users := make([]*User, 0)
	if err := DB.
		Preload("Team").
		Where("age = ?", 2).
		Find(&users).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	slice := AnyToAnySlice(users)

	assert.Equal(t, len(slice), 2)
	assert.Equal(t, len(slice[0].(*User).Team), 4)

}

func TestGORMByUtil(t *testing.T) {
	jinzhu1 := User{Name: "jinzhu1", Age: 2}
	jinzhu2 := User{Name: "jinzhu2", Age: 2}
	userA := User{Name: "A", ManagerID: &jinzhu1.ID}
	userB := User{Name: "B", ManagerID: &jinzhu1.ID}
	userC := User{Name: "C", ManagerID: &jinzhu1.ID}
	userD := User{Name: "D", ManagerID: &jinzhu1.ID}

	DB.Create(&jinzhu1)
	DB.Create(&jinzhu2)
	DB.Create(&userA)
	DB.Create(&userB)
	DB.Create(&userC)
	DB.Create(&userD)

	users := NewModelSlice(&User{}).Interface()

	if err := DB.
		Preload("Team").
		Where("age = ?", 2).
		Find(&users).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	slice := AnyToAnySlice(users)

	assert.Equal(t, len(slice), 2)
	assert.Equal(t, len(slice[0].(*User).Team), 4)
}
