package main

import (
	"reflect"
	"sort"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	var err error
	DB.Migrator().DropTable("order_0", "order_1", "users")
	err = DB.AutoMigrate(&User{})
	if err != nil {
		t.Error(err)
	}

	err = DB.Table("order_0").AutoMigrate(&Order{})
	if err != nil {
		t.Error(err)
	}

	// panic
	err = DB.Table("order_1").AutoMigrate(&Order{})
	if err != nil {
		t.Error(err)
	}

	tableList, _ := DB.Migrator().GetTables()

	targetList := []string{"order_0", "order_1", "users"}
	sort.Strings(targetList)
	sort.Strings(tableList)

	if !reflect.DeepEqual(tableList, targetList) {
		t.Errorf("table list %s", tableList)
	}
}
