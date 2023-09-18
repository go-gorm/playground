package main

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	ID              uint           `gorm:"column:id;not null;primary_key;AUTO_INCREMENT" json:"id"`
	ContentID       uint           `gorm:"column:content_id;default:null;index" json:"content_id"`
	AuthorID        uint           `gorm:"column:author_id;default:null;index" json:"author_id"`
	Content         string         `gorm:"column:content;default:null;type:text" json:"content"`
	ParentCommentID *uint          `gorm:"column:parent_comment_id;default:0" json:"parent_comment_id"`
	IP              uint           `gorm:"column:ip;default:0" json:"ip"`
	UserAgent       string         `gorm:"column:user_agent;default:null;type:varchar(1000)" json:"user_agent"`
	PushedAt        time.Time      `gorm:"column:pushed_at;default:null;type:datetime" json:"pushed_at"`
	AuthorName      string         `gorm:"column:author_name;default:null;type:varchar(32)" json:"author_name"`
	AuthorAvatarURL string         `gorm:"column:author_avatar_url;default:null;type:varchar(255)" json:"author_avatar_url"`
	AuthorURL       string         `gorm:"column:author_url;default:null;type:varchar(255)" json:"author_url"`
	BlogID          uint           `gorm:"column:blog_id;default:null;index " json:"blog_id"`
	Status          uint8          `gorm:"column:status;default:0;index " json:"status"`
	CreatedAt       time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP;type:datetime" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"column:updated_at;type:datetime" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at;index;type:datetime"  json:"deleted_at"`

	User User `gorm:"foreignKey:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
}
