package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
	"path/filepath"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}

func TestJoin(t *testing.T) {
	// When gorm.DB with `TablePrefix` config, the select column cannot match auto like `Model.field`. It's have to use
	// tool look like `Table` to make table with prefix.

	//adjust
	uTb := Table{}.Init("user")
	cTb := Table{}.Init("company")
	var users []User

	// uTb.Field()
	fields := uTb.Field("id")
	fields = append(fields, cTb.Field("name")...)

	DB.Select(fields).
		//Joins("Company", DB.Where(&Company{Name: "JC.inc"})).
		Joins("Company").
		Where(map[string]interface{}{"User.name": "jinzhu"}).
		Where(DB.Where(&Company{Name: "JC.inc"})).
		Find(&users)

	// failure
	DB.Select([]string{"User.id", "Company.Name"}).
		//Joins("Company", DB.Where(&Company{Name: "JC.inc"})).
		Joins("Company").
		Where(map[string]interface{}{"User.name": "jinzhu"}).
		Where(DB.Where(&Company{Name: "JC.inc"})).
		Find(&users)
}

func TestJoin2(t *testing.T) {
	var users []User
	db2 := getDbWithoutTablePrefix()
	db2.Select([]string{"User.id", "Company.Name"}).
		Joins("Company").
		Where(map[string]interface{}{"User.name": "jinzhu"}).
		Where(db2.Model(&Company{Name: "JC.inc"})).
		Find(&users)
}

func getDbWithoutTablePrefix() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(filepath.Join(os.TempDir(), "gorm.db")), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	return db
}
