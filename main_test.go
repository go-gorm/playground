package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

var DB *gorm.DB

func TestGORM(t *testing.T) {
	subQuery := DB.Model(Company{}).Where("id>?", 1)

	var testData []*User
	if err := DB.Model(User{}).Where("company_id>?", subQuery).Find(&testData); err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
