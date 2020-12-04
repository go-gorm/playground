package main

import "time"

type Form struct {
	ID       uint   `gorm:"column:id;not null;primary_key;AUTO_INCREMENT" json:"id"`
	UserID   uint   `gorm:"column:user_id;default:null;unique" json:"user_id"`
	UserType int    `gorm:"column:user_type;default:0;comment:'';type:tinyint(1)" json:"user_type"`
	Title    string `gorm:"column:title;default:null;comment:'';type:varchar(50)" json:"title"`

	User User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
}

type User struct {
	ID        uint      `gorm:"column:id;not null;primary_key;AUTO_INCREMENT" json:"id"`
	Name      string    `gorm:"column:name;default:null;comment:'';type:varchar(32)" json:"name"`
	ChannelID uint      `gorm:"column:channel_id;comment:'';default:null;" json:"channel_id"`
	SignTime  time.Time `gorm:"column:sign_time;default:null;comment:'';type:datetime" json:"sign_time"`

	Channel Channel `gorm:"foreignKey:ChannelID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"channel"`
}

type Channel struct {
	ID    uint   `gorm:"column:id;not null;primary_key;AUTO_INCREMENT" json:"id"`
	Title string `gorm:"column:title;default:null;comment:'';type:varchar(100)" json:"title"`
}
