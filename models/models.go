package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

type TestUser struct {
	ID        uint `gorm:"primarykey"`
	Name      string
	Email     string `gorm:"column:email;type:varchar(255);not null;index:idx_download_email_covering, option:INCLUDE (status, friendly)"`
	Status    string
	Friendly  string
	CreatedAt time.Time `gorm:"column:created_at;not null;index:idx_download_email_covering"`
}
