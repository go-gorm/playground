package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserMainID uint   `gorm:"column:user_main_id;not null;" json:"user_main_id"` // 用户主表ID
	Username   string `gorm:"column:username;not null;" json:"username"`         // 用户名

	UserExtInfo             *UserExt             `gorm:"foreignKey:UserID;" json:"user_ext_info"`              // 用户扩展信息，关联user_ext表
	UserAccountRelationInfo *UserAccountRelation `gorm:"foreignKey:UserID;" json:"user_account_relation_info"` // 用户账户关联信息，关联user_account_relation表
}

type UserExt struct {
	gorm.Model
	UserID uint `gorm:"column:user_id;not null;" json:"user_id"` // 用户ID

	UserInfo *User `gorm:"foreignKey:UserID;references:ID" json:"user_info"` // 用户信息，关联user表
}

type UserAccountRelation struct {
	gorm.Model

	UserID    uint `gorm:"column:user_id;not null;" json:"user_id"`       // 用户ID
	AccountID uint `gorm:"column:account_id;not null;" json:"account_id"` // 账户ID

	UserInfo    *User    `gorm:"foreignKey:UserID;references:ID" json:"user_info"`       // 用户信息，关联user表
	AccountInfo *Account `gorm:"foreignKey:AccountID;references:ID" json:"account_info"` // 账户信息，关联account表
}

type Account struct {
	gorm.Model
	CompanyID uint `gorm:"column:company_id;not null;" json:"company_id"` // 公司ID

	CompanyInfo *Company `gorm:"foreignKey:CompanyID;references:ID" json:"company_info"` // 公司信息，关联company表
}

type Company struct {
	gorm.Model
	Name string `gorm:"column:name;not null;" json:"name"` // 公司名称
}
