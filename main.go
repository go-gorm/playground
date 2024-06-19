package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	ID           uint           // Standard field for the primary key
	Name         string         // 一个常规字符串字段
	Email        *string        // 一个指向字符串的指针, allowing for null values
	Age          uint8          // 一个未签名的8位整数
	Birthday     *time.Time     // A pointer to time.Time, can be null
	MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
	ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
	CreatedAt    time.Time      // 创建时间（由GORM自动管理）
	UpdatedAt    time.Time      // 最后一次更新时间（由GORM自动管理）
}

func main() {
	dsn := "gorm:gorm@tcp(localhost:9910)/gorm"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		QueryFields: true,
		Logger:      logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(new(User))
	if err != nil {
		panic(err)
	}
	u := &User{
		Name: "test_name",
	}
	err = db.Create(u).Error
	if err != nil {
		panic(err)
	}
	j, err := json.MarshalIndent(&u, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(j))
	u = &User{
		ID:   u.ID,
		Name: "update_name",
	}
	err = db.Save(u).Error
	if err != nil {
		panic(err)
	}
	j, err = json.MarshalIndent(&u, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(j))
}
