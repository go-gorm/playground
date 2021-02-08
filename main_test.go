package main

import (
	"testing"

	"gorm.io/gorm/clause"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	// creat first line
	user := User{Name: "alice", Age: 18}
	user.ID = 1
	DB.Create(&user)

	// create duplicate line
	dupUser := User{Name: "alice", Age: 18}
	dupUser.ID = 1
	DB.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"age": 20,
		}),
		Where: clause.Where{Exprs: []clause.Expression{
			clause.Eq{
				Column: "name",
				Value:  "bob",
			},
		}},
	}).Create(&dupUser)

	// get result (except user.age == 18)
	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	t.Logf("\n%+v", result)
}
