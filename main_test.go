package main

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"testing"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type Instant time.Time

func (instant *Instant) Scan(value interface{}) error {
	tm := sql.NullTime{}
	if err := tm.Scan(value); err != nil {
		return err
	}

	if tm.Valid {
		*instant = Instant(tm.Time)
	}

	return nil
}

func (instant Instant) Value() (driver.Value, error) {
	return time.Time(instant), nil
}

func (instant Instant) GormDataType() string {
	return "timestamptz"
}

type MyModel struct {
	ID        uint           `gorm:"primarykey"`
	CreatedAt Instant        `gorm:"autocreatetime"`
	UpdatedAt Instant        `gorm:"autoupdatetime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type MyObject struct {
	MyModel
	V int `gorm:""`
}

func TestGORM(t *testing.T) {
	if err := DB.AutoMigrate(&MyObject{}); err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	DB.Config.NowFunc = func() time.Time {
		return time.Date(2022, time.July, 7, 0, 0, 0, 0, time.UTC)
	}

	obj := MyObject{V: 1}

	resultCreate := DB.Create(&obj)
	if resultCreate.Error != nil {
		t.Errorf("Failed, got error: %v", resultCreate.Error)
	}

	exp1 := resultCreate.Statement.Clauses["VALUES"].Expression.(clause.Values)
	fmt.Printf("Column 1: %s\n", exp1.Columns[0].Name)
	if _, ok := exp1.Values[0][0].(Instant); !ok {
		t.Error("First param should be an Instant")
	}

	fmt.Printf("Column 2: %s\n", exp1.Columns[1].Name)
	if _, ok := exp1.Values[0][1].(Instant); !ok {
		t.Error("Second param should be an Instant")
	}

	resultUpdate := DB.Save(&obj)
	if resultUpdate.Error != nil {
		t.Errorf("Failed, got error: %v", resultUpdate.Error)
	}

	exp2 := resultUpdate.Statement.Clauses["SET"].Expression.(clause.Set)
	fmt.Printf("Column 1: %s\n", exp2[0].Column.Name)
	if _, ok := exp2[0].Value.(Instant); !ok {
		t.Error("First param should be an Instant")
	}

	fmt.Printf("Column 2: %s\n", exp2[1].Column.Name)
	if _, ok := exp2[1].Value.(Instant); !ok {
		t.Error("Second param should be an Instant")
	}
}
