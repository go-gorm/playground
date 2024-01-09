package main

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strings"
)

type TestUser struct {
	Id               int    `json:"id"`
	Username         string `json:"username" gorm:"unique;index" validate:"max=12"`
	Password         string `json:"password" gorm:"not null;" validate:"min=8,max=20"`
	DisplayName      string `json:"display_name" gorm:"index" validate:"max=20"`
	Role             int    `json:"role" gorm:"type:int;default:1"`   // admin, common
	Status           int    `json:"status" gorm:"type:int;default:1"` // enabled, disabled
	Email            string `json:"email" gorm:"index" validate:"max=50"`
	GitHubId         string `json:"github_id" gorm:"column:github_id;index"`
	WeChatId         string `json:"wechat_id" gorm:"column:wechat_id;index"`
	VerificationCode string `json:"verification_code" gorm:"-:all"`                                    // this field is only for Email verification, don't save it to database!
	AccessToken      string `json:"access_token" gorm:"type:char(32);column:access_token;uniqueIndex"` // this token is for system management
	Quota            int    `json:"quota" gorm:"type:int;default:0"`
	UsedQuota        int    `json:"used_quota" gorm:"type:int;default:0;column:used_quota"` // used quota
	RequestCount     int    `json:"request_count" gorm:"type:int;default:0;"`               // request number
	Group            string `json:"group" gorm:"type:varchar(32);default:'default'"`
	AffCode          string `json:"aff_code" gorm:"type:varchar(32);column:aff_code;uniqueIndex"`
	InviterId        int    `json:"inviter_id" gorm:"type:int;column:inviter_id;index"`
}

func main() {
	dsn := "root:123456@tcp(localhost:13306)/chat?charset=utf8mb4&parseTime=True&loc=Local&timeout=10s"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true, // precompile SQL
		Logger:      logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println(err)
	}
	db = db.Set("gorm:table_options", "CHARSET=utf8mb4")
	err = db.AutoMigrate(&TestUser{})
	if err != nil {
		fmt.Println(err)
	}

	var user TestUser
	if err := db.Take(&user).Error; err != nil {
		rootUser := TestUser{
			Username:    "root",
			Password:    "$2a$10$OoNwis82Y7/yjN80Ces/FR1CZhK5YjOo7xn6pOTewobO",
			Role:        1,
			Status:      1,
			DisplayName: "Root User",
			AccessToken: strings.Replace(uuid.New().String(), "-", "", -1),
			Quota:       100000000,
		}
		err2 := db.Create(&rootUser).Error
		fmt.Println("-----------------------WHY???------------------------")
		fmt.Println("root user", err2)
	}

}
