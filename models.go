package main

type User struct {
	ID int64

	// users that are blocked from contacting this user
	BlockedUsers []*User `json:"-" gorm:"many2many:user_blocked_users;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// we don't want this field inside the database because it is a
	// computed field that gets the value based on a case statement
	// but we do need it on the model instance at times
	IsBlocked bool `gorm:"->;-:migration"`
}

func (u *User) Block(blockedUser *User) error {
	return DB.Model(&u).Association("BlockedUsers").Append(blockedUser)
}

type Chat struct {
	ID    int64
	Users []*User `json:"-" gorm:"many2many:user_chats;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (c *Chat) AddUser(user *User) error {
	return DB.Model(&c).Association("Users").Append(user)
}
