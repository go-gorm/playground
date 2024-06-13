package main

import (
	"time"
)

type Item struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement:true"`
	Name      string    `gorm:"unique"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
