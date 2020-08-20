package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type T01 struct {
	ID   int    `gorm:"autoIncrement"`
	Name string `gorm:"uniqueIndex"`
	T2s  []T02  `gorm:"many2many:t01_t02_bind; foreignKey:Name; joinForeignKey:Name1; references:Name; joinReferences:Name2; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
type T02 struct {
	ID   int    `gorm:"autoIncrement"`
	Name string `gorm:"uniqueIndex"`
}
func TestTest2(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "", log.LstdFlags),
			logger.Config{
				LogLevel: logger.Info,
			}),
	})
	checkError(err)
	
	err = db.AutoMigrate(&T01{}, &T02{})
	checkError(err)

	t1 := T01{Name: "1"}
	err = db.Create(&t1).Error
	checkError(err)
	
	err = db.Model(&t1).Association("T2s").Append([]T02{{Name: "1"}})
	checkError(err)
	err = db.Model(&t1).Association("T2s").Append([]T02{{Name: "1-1"}})
	checkError(err)
	
	t2s := make([]T02, 0, 2)
	err = db.Model(&t1).Association("T2s").Find(&t2s)
	checkError(err)
	
	b, _ := json.MarshalIndent(t2s, "", "    ")
	fmt.Printf("%s\n", b)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}
