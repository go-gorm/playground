package main

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SimpleFiles []map[string]any

// Value 实现 driver.Valuer 接口，将整个切片序列化为 JSON 数组
func (s SimpleFiles) Value() (driver.Value, error) {
	if s == nil {
		return nil, nil
	}
	return json.Marshal(s)
}

// Scan 实现 sql.Scanner 接口，将 JSON 数组解析到切片中
func (s *SimpleFiles) Scan(value interface{}) error {
	*s = make(SimpleFiles, 0) // 初始化为空切片
	if value == nil {
		return nil // 处理数据库的 NULL
	}

	b, ok := value.([]byte)
	if !ok {
		return nil
	}

	// 处理空字节切片（如空字符串）
	if len(b) == 0 {
		return nil
	}

	// 尝试解析 JSON
	var temp SimpleFiles
	if err := json.Unmarshal(b, &temp); err != nil {
		// 解析失败时返回空切片（如无效 JSON）
		return nil
	}

	*s = temp
	return nil
}

type Module struct {
	ID     uint         `gorm:"primarykey"`
	Images *SimpleFiles `gorm:"type:json" json:"images"`
}

func main() {

	dbDSN := "root:root@tcp(192.168.56.4:3306)/micro?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dbDSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	defer func() {
		db, _ := db.DB()
		db.Close()
	}()
	// insert into module (id,images) values (1,'[{"name":"1.jpg","url":"http://www.baidu.com"},{"name":"2.jpg","url":"http://www.baidu.com"}]')
	// insert into module (id) values (2)
	// insert into module (id) values (3, "")
	lists := []*Module{}
	db.Model(&Module{}).Find(&lists)

	// 1, 内容正确,
	// 2, images nil 不符合预期
	// 3, images [] 符合预期
	for _, item := range lists {
		b, _ := json.Marshal(item.Images)
		fmt.Printf("id: %d, images: %s\n", item.ID, string(b))
		// 1, [{"name":"1.jpg","url":"http://www.baidu.com"},{"name":"2.jpg","url":"http://www.baidu.com"}]
		// 2, null
		// 3, []
	}
}
