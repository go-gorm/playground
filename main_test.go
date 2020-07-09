package main

import (
	"testing"
	"encoding/json"
	"database/sql/driver"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type Map map[string]string
func (m Map) Value() (driver.Value, error) {
	b, err := json.Marshal(m)
	return string(b), err
}

func (m *Map) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &m)
}

type Test1 struct {
	Map map[string]string
}

type Test2 struct {
	Map Map
}

func TestGORM(t *testing.T) {
	t.Run("Regular Map", func(t *testing.T) {
		err := DB.AutoMigrate(&Test1{})
		if err != nil {
			t.Errorf("Failed, got error: %v", err)
		}
	})
	t.Run("Driver Map", func(t *testing.T) {
		err := DB.AutoMigrate(&Test2{})
		if err != nil {
			t.Errorf("Failed, got error: %v", err)
		}
	})
}
