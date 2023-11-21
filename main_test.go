package main

import (
	"database/sql"
	"testing"

	"gorm.io/gorm/clause"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	DB.Create(&Toy{Name: "1", OwnerID: "2", OwnerType: "3"})
	DB.Create(&Account{UserID: sql.NullInt64{Int64: 999}, Number: "666"})
	t.Log("create success")
	dbUpsert := DB.Clauses(clause.OnConflict{UpdateAll: true})
	dbUpsert.Create(&Toy{Name: "1", OwnerID: "2", OwnerType: "3"})
	dbUpsert.Create(&Account{UserID: sql.NullInt64{Int64: 999}, Number: "666"})
}
