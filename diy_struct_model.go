package main

import (
	"gorm.io/gorm"
	"time"
)

type Article struct {
	ID    uint   `gorm:"column:id;not null;primary_key;AUTO_INCREMENT" json:"id"`
	Title string `gorm:"column:title;not null;type:varchar(128)" json:"title"`

	CreatedAt time.Time  `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP;type:datetime" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`
	DeletedAt gorm.DeletedAt
}
