package main

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Logo     string        `json:"logo" gorm:"not null;type:varchar(50)"`
	Contents []ItemContent `json:"contents" gorm:"foreignKey:ItemID"`
}

type ItemContent struct {
	gorm.Model
	ItemID       uint   `json:"item_id" gorm:"not null"`
	Name         string `json:"name" gorm:"not null;type:varchar(50)"`
	LanguageCode string `json:"language_code" gorm:"not null;type:varchar(2)"`
}
