package main

import (
	_ "database/sql"
	_ "time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title   string
	Content string
	Threads []Thread `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"` // This wouldn't automagically work if soft-delete
}

type Thread struct {
	gorm.Model
	Title   string
	Content string
	PostID  uint
}

func (p *Post) BeforeDelete(db *gorm.DB) (err error) {
	tx := db.Delete(&p.Threads)
	return tx.Error
}
