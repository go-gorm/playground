package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	subQuery := DB.Model(Company{}).Select("id").Where("id > ?", 1)
	subQuery = subQuery.Where("name <> ?", "asdf")

	var testData []*User
	if err := DB.Model(User{}).Where("age>?", 18)Where("company_id > ?", subQuery).Find(&testData).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
