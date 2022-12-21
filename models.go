package main

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

type UpdateCount uint

type User struct {
	ID          uint `gorm:"primarykey"`
	Name        string
	UpdateCount UpdateCount `gorm:"default:0"`
}

func (u UpdateCount) UpdateClauses(field *schema.Field) []clause.Interface {
	return []clause.Interface{updateCountClause{Field: field}}
}

func (u UpdateCount) CreateClauses(field *schema.Field) []clause.Interface {
	return []clause.Interface{updateCountClause{Field: field}}
}

type updateCountClause struct {
	Field *schema.Field
}

func (u updateCountClause) Name() string {
	return ""
}

func (u updateCountClause) Build(builder clause.Builder) {
}

func (u updateCountClause) MergeClause(c *clause.Clause) {
}

func (u updateCountClause) ModifyStatement(stmt *gorm.Statement) {
	if stmt.SQL.Len() == 0 && !stmt.Statement.Unscoped {
		var count = 0
		v, zero := u.Field.ValueOf(stmt.Context, stmt.ReflectValue)
		if !zero {
			count = int(v.(UpdateCount))
		}
		stmt.AddClause(clause.Set{{Column: clause.Column{Name: u.Field.DBName}, Value: count}})
		stmt.SetColumn(u.Field.DBName, count, true)
	}
}
