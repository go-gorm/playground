package main

type Role string

const (
	VIEWER Role = "VIEWER"
	EDITOR Role = "EDITOR"
	ADMIN  Role = "ADMIN"
)

// User has a UserID field which is of type string.
type User struct {
	UserID string `gorm:"column:user_id;primaryKey" json:"user_id"`
}

// UserPermission Binds a User to Another Table through a Role.
// for simplicity the other table has been ommited.
type UserPermission struct {
	ID     uint   `gorm:"primaryKey;autoIncrement" json:"-"`
	UserID string `gorm:"type:text;column:user_id;not null;index:idx_tenant_tech;constraint:OnDelete:CASCADE" json:"user_id"`
	User   User   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Role   Role   `gorm:"column:role;not null" json:"role"`
}
