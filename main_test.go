package main

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"strconv"
	"testing"
)

type Model struct {
	ID       uint64 `gorm:"primaryKey;autoIncrement"`
	Title    string `gorm:"size:100;not null"`
	TypeEnum int    `gorm:"-"`
	Type     string `gorm:"size:100"`
}

func (m *Model) BeforeSave(db *gorm.DB) error {
	if m.TypeEnum == 0 {
		fmt.Println(*m) // see the console output, its completely empty
		return errors.New("illegal type")
	}
	db.Statement.SetColumn("Type", strconv.Itoa(m.TypeEnum))
	return nil
}

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	err := DB.AutoMigrate(&Model{})
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	m := &Model{
		Title:    "New",
		TypeEnum: 1,
	}

	// Create the Model
	if err = DB.Create(m).Error; err != nil {
		t.FailNow()
	}

	// Update the Model
	m.TypeEnum = 3

	// Save the Model
	err = DB.Model(&Model{}).Where("id = ?", m.ID).Updates(m).Error
	if err != nil { // There is an "illegal type" error, because the model in the callback is empty
		t.Fatal(err)
	}
}
