package main

import (
	"testing"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func TestGenericsJoins(t *testing.T) {
	dbg := DB.Session(&gorm.Session{DryRun: true})
	db := gorm.G[User](dbg)

	q := db.Joins(clause.LeftJoin.AssociationFrom("Company", gorm.G[Company](dbg)).As("t"),
		func(j gorm.JoinBuilder, joinTable clause.Table, curTable clause.Table) error {
			j.Where("?.\"name\" = ?", joinTable, "GenericsCompany")
			return nil
		},
	).Where(map[string]any{"name": "GenericsJoins_2"})

	stmt := &gorm.Statement{DB: dbg}
	q.Build(stmt)

	sql := stmt.SQL.String()
	t.Logf("GENERATED SQL:\n%s", sql)
}
