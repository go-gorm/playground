package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	ids := make([]int, 0, 1000)
	users := make([]User, 0, 1000)
	for i := 1; i <= 1000; i++ {
		ids = append(ids, i)
		users = append(users, User{CompanyID: &i, Age: 10})
	}
	DB.Create(&users)
	idsLen := len(ids)
	start := 0
	for start < idsLen {
		end := start + 100
		if end > idsLen {
			end = idsLen
		}
		err := DB.Model(&User{}).Where("company_id in (?)", ids[start:end]).Update("age", 18).Error
		if err != nil {
			t.Errorf("Failed, got error: %v", err)
		}
		start = end
	}
}
