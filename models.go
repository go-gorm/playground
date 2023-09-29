package main

import (
	"database/sql"
	"gorm.io/gorm"
)

type TestTable struct {
	gorm.Model
	UserID sql.NullInt64
	Number string
}
