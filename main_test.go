package main

import (
	"encoding/hex"
	"github.com/google/uuid"
	"github.com/tzq0301/gorms"
	"strings"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	db, _ := gorms.InMemorySqliteWithInitSQL(`
		CREATE TABLE objects (
  			id BINARY(16)       
		);
	`)

	id1, _ := hex.DecodeString(strings.ReplaceAll(uuid.New().String(), "-", ""))
	id2, _ := hex.DecodeString(strings.ReplaceAll(uuid.New().String(), "-", ""))
	id3, _ := hex.DecodeString(strings.ReplaceAll(uuid.New().String(), "-", ""))

	db.Debug().Create(&Object{ID: uuid.UUID(id1)})
	db.Debug().Create(&Object{ID: uuid.UUID(id2)})
	db.Debug().Create(&Object{ID: uuid.UUID(id3)})

	ids := [][]byte{
		id1,
		id2,
		id3,
	}

	//var obj []*Obj
	query := `
	select * from objects where id in (?, ?, ?)
	`
	err := db.Exec(query, ids)
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

}

type Object struct {
	ID uuid.UUID
}
