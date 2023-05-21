package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type Device struct {
	ID      int `gorm:"primarykey"`
	Class   string
	Details []DeviceDetail `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type DeviceDetail struct {
	DeviceID int    `gorm:"index:idx_unique"`
	Key      string `gorm:"index:idx_unique"`
	Value    string
}

func TestGORM(t *testing.T) {
	dev := Device{
		Class: "foo",
		Details: []DeviceDetail{{
			Key:   "key",
			Value: "val",
		}},
	}

	if err := DB.AutoMigrate(new(Device), new(DeviceDetail)); err != nil {
		t.Error("migration", err)
	}

	tx := DB.Create(&dev)
	if err := tx.Error; err != nil {
		t.Error("create", err)
	}

	var row Device
	if err := tx.First(&row).Error; err != nil {
		t.Error("first", err)
	}

	if err := DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(DeviceDetail{DeviceID: row.ID}).Error; err != nil {
			return err
		}

		return tx.Delete(Device{ID: row.ID}).Error
	}); err != nil {
		t.Error("delete", err)
	}
}
