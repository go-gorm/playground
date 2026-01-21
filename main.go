package main

import (
	"log"

	"github.com/gogf/gf/v2/util/grand"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Route struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt int64
	UpdatedAt int64
	Name      string
}
type Update struct {
	Name      string
	UpdatedAt int64
}
type NoUpdate struct {
	Name string
}

func main() {
	// 连接到 MySQL 数据库
	dsn := "root:xxxx@tcp(xxxx:3306)/xxxx?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	var route = &Route{}
	// 使用 Updates 方法更新 Route
	//[7.219ms] [rows:1] UPDATE `routes` SET `updated_at`=1728225917,`name`='KgqHaJnfID' WHERE `id` = 5
	result := db.Model(&route).Debug().Where("id", 5).Updates(Route{Name: grand.Letters(10)})
	if result.Error != nil {
		log.Fatalf("failed to update route: %v", result.Error)
	}
	//[8.108ms] [rows:1] UPDATE `routes` SET `updated_at`=1728225917,`name`='zPrQbQopwm' WHERE `id` = 5
	result = db.Model(&route).Debug().Where("id", 5).Updates(Update{Name: grand.Letters(10)})
	if result.Error != nil {
		log.Fatalf("failed to update route: %v", result.Error)
	}

	//[7.656ms] [rows:1] UPDATE `routes` SET `name`='NkGiUbrINK' WHERE `id` = 5
	result = db.Model(&route).Debug().Where("id", 5).Updates(NoUpdate{Name: grand.Letters(10)})
	if result.Error != nil {
		log.Fatalf("failed to update route: %v", result.Error)
	}
}
