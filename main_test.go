package main

import (
	"strings"
	"testing"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {

	var result User

	newDB := DB.Session(&gorm.Session{NewDB: true, DryRun: true})

	newDB = newDB.Table("users")
	newDB.Clauses(
		clause.From{
			Tables: []clause.Table{{Name: "users"}},
			Joins: []clause.Join{
				{
					Table: clause.Table{Name: "companies", Raw: false},
					ON: clause.Where{
						Exprs: []clause.Expression{
							clause.Eq{
								Column: clause.Column{
									Table: "users",
									Name:  "company_id",
								},
								Value: clause.Column{
									Table: "companies",
									Name:  "id",
								},
							},
						},
					},
				},
			},
		},
	)

	newDB.Joins("inner join rgs on rgs.id = user.id")

	stmt := newDB.First(&result).Statement
	str := stmt.SQL.String()

	if !strings.Contains(str, "rgs.id = user.id") {
		t.Errorf("The second join condition is over written instead of combining")
	}

	if !strings.Contains(str, "`users`.`company_id` = `companies`.`id`") && !strings.Contains(str, "\"users\".\"company_id\" = \"companies\".\"id\"") {
		t.Errorf("The first join condition is over written instead of combining")
	}

}
