package main

import (
	"log"

	"gorm.io/driver/sqlite" // Any dialector
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	var dbModel TestStruct

	// there's a panic in Scan and a potential hang in gorm
	err = db.Raw("SELECT 'test' as payload").First(&dbModel).Error
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Result: %+v\n", dbModel)
}

type TestStruct struct {
	Payload OwnerType `gorm:"column:payload"`
}

type OwnerType string

func (s *OwnerType) Scan(value interface{}) error {
	data, ok := value.(string)
	if !ok {
		*s = ""
		return nil
	}

	// panic
	data = data[:-len(data)]

	*s = OwnerType(data)

	return nil
}
