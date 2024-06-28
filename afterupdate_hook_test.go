package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

type Product6 struct {
	ID       uint `gorm:"primarykey"`
	Name     *string
	UniqueId *uuid.UUID `gorm:"type:varchar(36);unique" json:",omitempty"`
}

// Hold copy of the values passed to the AfterUpdate() for subsequent check
var hookName *string
var hookUuid *uuid.UUID

func (p *Product6) AfterUpdate(*gorm.DB) error {

	hookName = p.Name
	hookUuid = p.UniqueId

	name := "nil"
	if p.Name != nil {
		name = *p.Name
	}
	uniqueId := "nil"
	if p.UniqueId != nil {
		uniqueId = p.UniqueId.String()
	}
	fmt.Printf("After update p.UniqueId %v p.Name %v\n", uniqueId, name)
	return nil
}

func StringPtr(s string) *string {
	return &s
}
func UuidPtr(u uuid.UUID) *uuid.UUID {
	return &u
}

func TestAfterUpdateHook(t *testing.T) {
	DB.Migrator().DropTable(&Product6{})
	DB.Debug().AutoMigrate(&Product6{})

	productName := StringPtr("product name")
	productUuid := UuidPtr(uuid.MustParse("5f003fc3-f806-42a3-ae4e-4d4335581de0"))
	p := Product6{Name: productName, UniqueId: productUuid}
	DB.Model(&Product6{}).Create(&p)

	p.Name = nil
	p.UniqueId = nil
	err := DB.Model(&p).Updates(map[string]interface{}{"name": p.Name, "unique_id": p.UniqueId}).Error
	assert.Nil(t, err)
	assert.Nil(t, hookName)
	assert.Nil(t, hookUuid)

	productName = StringPtr("new name")
	productUuid = UuidPtr(uuid.MustParse("415d9c8a-8742-4fe6-a6cd-8062a121eeb4"))
	err = DB.Model(&p).Updates(map[string]interface{}{"name": productName, "unique_id": productUuid}).Error
	assert.Nil(t, err)
	assert.Equal(t, hookName, productName)
	assert.Equal(t, hookUuid, productUuid)

	productName = nil
	productUuid = nil
	err = DB.Model(&p).Updates(map[string]interface{}{"name": productName, "unique_id": productUuid}).Error
	assert.Nil(t, err)
	assert.Nil(t, hookName)
	assert.Nil(t, hookUuid)
}
