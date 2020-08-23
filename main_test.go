package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type T01 struct {
	ID   int    `gorm:"autoIncrement"`
	Name string `gorm:"uniqueIndex"`
	T02s []T02  `gorm:"many2many:t01_t02_bind; foreignKey:Name; joinForeignKey:Name1; references:Name; joinReferences:Name2; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
type T02 struct {
	ID   int    `gorm:"autoIncrement"`
	Name string `gorm:"uniqueIndex"`
}

func TestTest2(t *testing.T) {
	db, err := gorm.Open(postgres.Open("database=postgres"), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "", log.LstdFlags),
			logger.Config{
				LogLevel: logger.Info,
				Colorful: true,
			}),
	})
	checkError(err)
	err = db.AutoMigrate(&T01{}, &T02{})
	checkError(err)
	t1 := T01{Name: "1"}
	err = db.Create(&t1).Error
	checkError(err)
	tt := t1

	err = db.Model(&t1).Association("T02s").Append([]T02{{Name: "1-1"}, {Name: "1-2"}})
	checkError(err)
	prettyPrint(t1)

	t1 = tt
	err = db.Model(&t1).Association("T02s").Replace([]T02{{Name: "1-1"}})
	checkError(err)
	prettyPrint(t1)

	t1 = tt
	err = db.Model(&t1).Association("T02s").Replace([]T02{{Name: "1-1"}, {Name: "1-2"}, {Name: "1-3"}})
	checkError(err)
	prettyPrint(t1)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

func prettyPrint(v interface{}) {
	b, _ := json.MarshalIndent(v, "", "    ")
	fmt.Printf("%s\n", b)
}
