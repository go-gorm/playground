package main

import (
	"testing"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

// Mock out the only relevant part of the dialector
type QuoteDialector struct {}
func (_ *QuoteDialector) Name() string { return "" }
func (_ *QuoteDialector) Initialize(_ *gorm.DB) error { return nil }
func (_ *QuoteDialector) Migrator(_ *gorm.DB) gorm.Migrator { return nil }
func (_ *QuoteDialector) DataTypeOf(_ *schema.Field) string { return "" }
func (_ *QuoteDialector) DefaultValueOf(_ *schema.Field) clause.Expression { return nil }
func (_ *QuoteDialector) BindVarTo(_ clause.Writer, _ *gorm.Statement, _ interface{}) {}
func (_ *QuoteDialector) QuoteTo(w clause.Writer, s string) {
	w.WriteString("\"")
	w.WriteString(s)
	w.WriteString("\"")
}
func (_ *QuoteDialector) Explain(sql string, vars ...interface{}) string { return "" }

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	stmt := &gorm.Statement{DB: &gorm.DB{Config: &gorm.Config{Dialector: &QuoteDialector{}}}}
	stmt.QuoteTo(&stmt.SQL, clause.Table{Name: "table", Alias: "alias"})
	expected := `"table" AS "alias"`
	if stmt.SQL.String() != expected {
		t.Errorf("Failed, generated SQL is not equal, expects `%v`, but got `%v`", expected, stmt.SQL.String())
	}
}
