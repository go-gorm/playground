package main

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Documents []*Document // has-many
}
type Document struct {
	gorm.Model
	User             *User // belongs-to
	UserID           uint
	Name             string
	DocumentFulltext *DocumentFulltext // has-one
}
type DocumentFulltext struct {
	gorm.Model
	DocumentID uint
	Document   *Document // belongs-to
	Name       string
}

type DocumentListEntry struct {
	Document       `gorm:"embedded"`
	FulltextExists bool
}
