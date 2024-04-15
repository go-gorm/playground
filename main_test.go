package main

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/k0kubun/pp/v3"
	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type Example struct {
	CreatedAt *time.Time      `json:"-"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"-"`
	UpdatedAt *time.Time      `json:"UpdatedAt,omitempty,format:unix"`
	UUID      uuid.UUID       `gorm:"primarykey;type:uuid" json:"UUID"`
	Data      string
}

type Examples struct {
	CreatedAt *time.Time      `json:"-"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"-"`
	UpdatedAt *time.Time      `json:"UpdatedAt,omitempty,format:unix"`
	UUID      uuid.UUID       `gorm:"primarykey;type:uuid;default:gen_random_uuid()" json:"UUID"`
	Data      string
}

type Example2 struct {
	CreatedAt *time.Time      `json:"-"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"-"`
	UpdatedAt *time.Time      `json:"UpdatedAt,omitempty,format:unix"`
	UUID      uuid.UUID       `gorm:"primarykey;type:uuid;default:gen_random_uuid()" json:"UUID"`
	Data      string
}

func TestGORM(t *testing.T) {

	DB.Migrator().DropTable(&Example2{})
	DB.Migrator().DropTable(&Example{})
	DB.Migrator().DropTable(&Examples{})

	//validate check for default as functional
	err := DB.AutoMigrate(&Example2{})
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	dat, err := DB.Migrator().ColumnTypes(&Example2{})
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	for _, v := range dat {
		if v.Name() == "uuid" {
			val, ok := v.DefaultValue()
			if !ok {
				t.Errorf("missing default value for example2")
			}
			if val != "gen_random_uuid()" {
				t.Errorf("incorrect default value for example2")
			}
		}
	}

	pp.Println(dat)

	//create example
	err = DB.AutoMigrate(&Example{})
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	dat, err = DB.Migrator().ColumnTypes(&Example{})
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	for _, v := range dat {
		if v.Name() == "uuid" {
			_, ok := v.DefaultValue()
			if ok {
				t.Errorf("unexpected default value for uuid")
			}
		}
	}

	//migrate example testing for default not existing where expected.
	err = DB.AutoMigrate(&Examples{})
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	dat, err = DB.Migrator().ColumnTypes(&Examples{})
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	for _, v := range dat {
		if v.Name() == "uuid" {
			val, ok := v.DefaultValue()
			if !ok {
				t.Errorf("missing default value for example2")
			}
			if val != "gen_random_uuid()" {
				t.Errorf("incorrect default value for example2")
			}
		}
	}

}
