package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type S struct {
	ID   int    `gorm:"autoIncrement"`
	Name string `gorm:""`
	Fs   []F    `gorm:"many2many:s_f_bind; foreignKey:ID; joinForeignKey:SID; references:ID; joinReferences:FID"`
}
type F struct {
	ID   int    `gorm:"autoIncrement"`
	Name string `gorm:""`
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), nil)
	checkError(err)
	err = db.AutoMigrate(&S{})
	checkError(err)
	err = db.AutoMigrate(&F{})
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
