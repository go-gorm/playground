package main

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

type HasTenant sql.NullString

// Scan implements the Scanner interface.
func (t *HasTenant) Scan(value interface{}) error {
	return (*sql.NullString)(t).Scan(value)
}

// Value implements the driver Valuer interface.
func (t HasTenant) Value() (driver.Value, error) {
	if !t.Valid {
		return nil, nil
	}
	return t.String, nil
}

func (HasTenant) QueryClauses(f *schema.Field) []clause.Interface {
	return []clause.Interface{HasTenantQueryClause{Field: f}}
}

type HasTenantQueryClause struct {
	Field *schema.Field
}

func (sd HasTenantQueryClause) Name() string {
	return ""
}

func (sd HasTenantQueryClause) Build(clause.Builder) {
}

func (sd HasTenantQueryClause) MergeClause(*clause.Clause) {
}

func (sd HasTenantQueryClause) ModifyStatement(stmt *gorm.Statement) {
	if stmt.Context==nil{
		fmt.Println("modify")
		panic("context nil")
	}
}