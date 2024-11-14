package main

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// User has one `Account` (has one), many `Pets` (has many) and `Toys` (has many - polymorphic)
// He works in a Company (belongs to), he has a Manager (belongs to - single-table), and also managed a Team (has many - single-table)
// He speaks many languages (many to many) and has many friends (many to many - single-table)
// His pet also has one Toy (has one - polymorphic)
type User struct {
	ID        uint `gorm:"primarykey"`
	Name      string
	CreatedAt Time `gorm:"not null"`
	UpdatedAt Time `gorm:"not null"`
}

type Account struct {
	gorm.Model
	UserID sql.NullInt64
	Number string
}

type Pet struct {
	gorm.Model
	UserID *uint
	Name   string
	Toy    Toy `gorm:"polymorphic:Owner;"`
}

type Toy struct {
	gorm.Model
	Name      string
	OwnerID   string
	OwnerType string
}

type Company struct {
	ID   int
	Name string
}

type Language struct {
	Code string `gorm:"primarykey"`
	Name string
}

type Time struct {
	date time.Time
}

func NewTime(year, month, day, hour, min, sec int) Time {
	return Time{
		date: time.Date(year, time.Month(month), day, hour, min, sec, 0, time.UTC),
	}
}

// GormDataType returns gorm common data type. This type is used for the field's column type.
func (Time) GormDataType() string {
	return "time"
}

// GormDBDataType returns gorm DB data type based on the current using database.
func (Time) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "mysql":
		return "TIME"
	case "postgres":
		return "TIME"
	case "sqlserver":
		return "TIME"
	case "sqlite":
		return "TEXT"
	default:
		return ""
	}
}

// Scan implements sql.Scanner interface and scans value into Time,
func (t *Time) Scan(src interface{}) error {
	switch v := src.(type) {
	case []byte:
		return t.setFromString(string(v))
	case string:
		return t.setFromString(v)
	case time.Time:
		t.setFromTime(v)
	default:
		return fmt.Errorf("failed to scan value: %v", v)
	}

	return nil
}

func (t *Time) setFromString(str string) error {
	time, err := time.Parse(time.DateTime, str)
	if err != nil {
		return err
	}
	t.date = time
	return nil
}

func (t *Time) setFromTime(src time.Time) {
	*t = NewTime(src.Year(), int(src.Month()), src.Day(), src.Hour(), src.Minute(), src.Second())
}

// Value implements driver.Valuer interface and returns string format of Time.
func (t Time) Value() (driver.Value, error) {
	return t.date.Format(time.DateTime), nil
}
