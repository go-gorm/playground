package main

import (
	"context"
	"database/sql"
	"reflect"
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

func TestConnectionsLeak(t *testing.T) {

	db := DB.Table("non_existent")

	db.WithContext(context.Background()).FirstOrCreate(&User{Name: "foo"})
	db.WithContext(context.Background()).FirstOrCreate(&User{Name: "foo"})
	db.WithContext(context.Background()).FirstOrCreate(&User{Name: "foo"})

	connPool := db.ConnPool.(*sql.DB)
	v := reflect.ValueOf(connPool).Elem()
	f := v.FieldByName("numOpen")

	if f.Int() > 0 {
		t.Errorf("Expected no open connections but found %d", f.Int())
	}

}
