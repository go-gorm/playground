package main

import (
	"context"
	"strings"
	"testing"

	"gorm.io/gorm"
	q "gorm.io/playground/dal/query"
)

func GetSQLStatementInScope(callback func()) string {
	var sql string
	DB.Callback().Update().Replace("gorm:raw_sql", func(db *gorm.DB) {
		sql = db.Statement.SQL.String()
	})
	callback()

	return sql
}

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	name := "jinzhu"
	user := User{Name: name}

	DB.Create(&user)

	query := q.Use(DB)
	ctx := context.Background()

	expect := GetSQLStatementInScope(func() {
		query.WriteDB().Transaction(func(tx *q.Query) error {
			user := query.User
			user.WithContext(ctx).Where(user.Name.Eq(name)).UpdateSimple(user.Name.Value("test2"))
			return nil
		})
	})

	actual := GetSQLStatementInScope(func() {
		query.WriteDB().Transaction(func(tx *q.Query) error {
			user := tx.User
			user.WithContext(ctx).Where(user.Name.Eq(name)).UpdateSimple(user.Name.Value("test2"))
			return nil
		})
	})

	// not equals
	if expect != actual {
		t.Errorf("NOT EQUALS: expect:%s, actual: %s", expect, actual)
	}

	// not update updated_at
	if !strings.Contains(actual, "updated_at") {
		t.Errorf("`updated_at` IS NOT TOUCHED IN QUERY: %s", actual)
	}
}
