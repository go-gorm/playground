package main

import (
)

type User struct {
	Name      string `gorm:"uniqueIndex;not null;default:"`
}
