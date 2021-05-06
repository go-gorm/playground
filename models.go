package main

import (
	"time"
)

type Model struct {
	PK        int64 ` gorm:"primaryKey"`
	ID        string
	CreatedBy *int64
	Created   *Model `gorm:"->;foreignKey:CreatedBy"`
	CreatedAt time.Time
}

type Account struct {
	Model
}
