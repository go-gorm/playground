package main

import (
	"gorm.io/gorm/clause"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: mysql

//func TestGORM(t *testing.T) {
//	user := User{Name: "jinzhu"}
//
//	DB.Create(&user)
//
//	var result User
//	if err := DB.First(&result, user.ID).Error; err != nil {
//		t.Errorf("Failed, got error: %v", err)
//	}
//}

type MyRecord struct {
	Id int64  `gorm:"<-:false"`          // disable write permission
	A  string `gorm:"index:uq_a,unique"` // unique index
	B  string
}

func TestMyRecord(t *testing.T) {
	rec := MyRecord{Id: 1002, A: "foo", B: "bar"}
	DB.Create(&rec)

	var result MyRecord
	// the record has primary key id=1, since we disable write permission for column `id`
	if err := DB.Model(&MyRecord{}).Where("id = ?", 1).First(&result).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	// we do update on unique index conflict, this time column `id` set successfully.
	DB.Clauses(clause.OnConflict{DoUpdates: clause.Assignments(map[string]interface{}{"id": rec.Id})}).Create(&rec)
	if err := DB.Model(&MyRecord{}).Where("id = ?", 1).First(&result).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
