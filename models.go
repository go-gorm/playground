package main

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint64    `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type File struct {
	BaseModel
	Name      string `json:"name,omitempty" gorm:"type:text"`
	Type      string `json:"type,omitempty" gorm:"type:varchar(50)"`
	SizeByte  uint64 `json:"size_byte,omitempty"`
	DataURL   string `json:"data_url" gorm:"type:text"`
	CommentID uint64 `json:"comment_id" gorm:"not null"`
}

// This is not call, when delete Comment with associated file
func (f *File) BeforeDelete(ct *gorm.DB) (err error) {
	fmt.Println("BeforeDelete file Called:", f, f.Name)
	return
}

type Comment struct {
	BaseModel
	Description string `json:"description" gorm:"type:text"`
	PinID       uint64 `json:"pin_id" gorm:"not null"`
	Files       []File `gorm:"constraint:OnDelete:CASCADE"`
}
