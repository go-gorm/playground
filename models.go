package main

import (
	"database/sql/driver"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Name  string
	State CustomType
}

type CustomType struct {
	Machine *Machine
}

type Machine string

func NewCustomType(state Machine) CustomType {
	return CustomType{&state}
}

// Implement scanner and value interfaces as described in tests https://gorm.io/docs/data_types.html#Scanner-x2F-Valuer
// https://github.com/go-gorm/gorm/blob/master/tests/scanner_valuer_test.go

func (sm *CustomType) Scan(value interface{}) error {
	log.Printf("Scanning state %v", value)
	switch vt := value.(type) {
	case string:
		state := vt
		*sm = NewCustomType(Machine(state))
	}
	return nil
}

func (sm CustomType) Value() (driver.Value, error) {
	log.Printf("Value of state %v", sm)
	return &sm.Machine, nil
}
