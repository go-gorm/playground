package main

import (
	"testing"
	"log"
	"github.com/fatih/structs"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu", Age: 33, Active:true}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	log.Println(result)
	data:=structs.Map(result)
	var updatedData User
	if err:=DB.Table("users").Where("id = ?",1).Omit("Name").Updates(data).Scan(&updatedData); err!=nil{
		t.Errorf("Failed, got error: %v",err)
	}
	log.Println(updatedData)
}
