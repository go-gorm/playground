package main

type User struct {
	ID     string      `gorm:"column:id;type:varchar(36);primary_key" json:"id"`
	Name   string      `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Avatar *UserAvatar `gorm:"foreignKey:UserID;constraint:OnDelete:SET NULL"`
}

type UserAvatar struct {
	ID     string `gorm:"column:id;type:varchar(36);primary_key" json:"id"`
	UserID string `gorm:"column:user_id;type:varchar(36);uniqueIndex" json:"userID"`
	Size   int64  `gorm:"column:size;type:bigint" json:"size"`
}

type userRoomAvatar struct {
	User
	Avatar *UserAvatar `gorm:"embedded;embeddedPrefix:avatar_"`
	RoomID string
}

type usersRoomAvatars []userRoomAvatar
