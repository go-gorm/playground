package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// "encoding/json"

func FindItems[T any](createDB /* func()  */*gorm.DB, limit int, page int, model *T, query map[string]any, order string, direction string) ([]map[string]any, error) {
	db := createDB/* () */
	// sqlDB, err := db.DB()
	// if err != nil {
	// 	panic(err)
	// }
	// defer sqlDB.Close()
	if direction == "desc" {
		db = db.Order(order + " " + "DESC")
	} else {
		db = db.Order(order + " " + "ASC")
	}
	var items = []map[string]any{}
	result := db.Model(model).Omit("deleted_at").Where(query).Limit(limit).Offset(page * limit).Find(&items)

	return items, result.Error
}

type DeletedAt = gorm.DeletedAt
type ToDoItem struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt DeletedAt `gorm:"index" json:"-"  `
	Content   string    `gorm:"not null" json:"content"  form:"content"`

	Completed bool `gorm:"not null;index" json:"completed" form:"completed" `

	ID uint `gorm:"primarykey" json:"id" form:"id"`

	Author string `gorm:"not null;index" json:"author" form:"author"`
}

func main() {
	dsn := "todolist:todolist@tcp(127.0.0.1:3306)/todolist?charset=utf8mb4&parseTime=True&loc=Local"
	createDB := ConnectDatabase(dsn, &ToDoItem{}, "to_do_items", true)
	for i := 0; i < 5; i++ {
		res, err := FindItems(createDB, 10, 0, &ToDoItem{}, map[string]any{}, "id", "desc")

		if err != nil {
			panic(err)
		}

		fmt.Printf("%v", len(res))
	}
	for i := 0; i < 5; i++ {
		res, err := FindItems(createDB, 10, 0, &ToDoItem{}, map[string]any{}, "completed", "asc")

		if err != nil {
			panic(err)
		}

		fmt.Printf("%v", len(res))
	}
}

/* func CloseDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	return sqlDB.Close()
}
 */
func ConnectDatabase[T any](dsn string, model *T, TableName string, debug bool) /* func() */ *gorm.DB {

	var createDB = func() *gorm.DB {
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		db = db.Table(TableName)
		if debug {
			db = db.Debug()
		}

		db = db.Model(model)
		sqlDB, err := db.DB()
		if err != nil {
			panic(err)
		}
		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		sqlDB.SetMaxIdleConns(10)

		// SetMaxOpenConns sets the maximum number of open connections to the database.
		sqlDB.SetMaxOpenConns(100)
		return db
	}

	db := createDB()
	// defer CloseDB(db)
	err := db.AutoMigrate(model)
	if err != nil {
		panic(err)
	}

	return createDB()
	// return createDB
}
