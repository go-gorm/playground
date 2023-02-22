package main

import "gorm.io/gorm"

func main() {

	gdb, _ := gorm.Open(sqlserver.Open("db uri"))

	_ = gdb.AutoMigrate(&User{}) // this step does not create the conditional index as expected in MS SQL
}
