package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	subQuery := DB.Model(Company{}).Select("id").Where("id>? AND name<>?", 1, "asdf")

	var testData []*User
	if err := DB.Model(User{}).Where("company_id>?", subQuery).Find(&testData).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
