package main

import (
	"testing"

	"gorm.io/gorm/clause"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	u1 := User{Name: "A"}
	if err := DB.Create(&u1).Error; err != nil {
		panic(err)
	}

	// set A as B's manager.
	if err := DB.Create(&User{Name: "B", ManagerID: &u1.ID}).Error; err != nil {
		panic(err)
	}

	var user = User{}
	if err := DB.Where("id", u1.ID).
		Preload("Team.Manager").
		Preload(clause.Associations).
		Find(&user).
		Error; err != nil {
		panic(err)
	}

	if (user.Team[0].Manager) == nil {
		// gorm.io/gorm v1.20.12
		t.Fatal("expect Manager A")
	}

	// Result:
	// gorm.io/gorm v1.20.12: FAIL: TestGORM (0.02s). ❌
	// gorm.io/gorm v1.20.9: PASS: TestGORM (0.02s). ✅
}
