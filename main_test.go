package main

import (
	"gorm.io/datatypes"
	"reflect"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type Param struct {
	ID          int
	DisplayName string
	Config      datatypes.JSON
}

func TestGORM(t *testing.T) {

	DB.Exec("DROP TABLE IF EXISTS `params`")

	DB.Exec("CREATE TABLE `params` (`id` tinyint NOT NULL, `display_name` varchar(64) NOT NULL, `config` json DEFAULT NULL, PRIMARY KEY (`id`))")

	DB.Exec("INSERT INTO `params` (`id`,`display_name`,`config`) VALUES (1,'Test1','{\"val1\": 987654, \"val2\": \"abcdefghijklmnopqrstuvwxyz\"}');")

	DB.Exec("INSERT INTO `params` (`id`,`display_name`,`config`) VALUES (2,'TEST_JSON','{\"param1\": 1234, \"param2\": \"test\"}')")

	fetchId := 2

	cmp1 := Param{
		ID:          fetchId,
		DisplayName: "TEST_JSON",
		Config:      datatypes.JSON("{\"param1\": 1234, \"param2\": \"test\"}"),
	}

	var retSingle1 Param

	err := DB.Where("id = ?", fetchId).Take(&retSingle1).Error
	if err != nil || reflect.DeepEqual(retSingle1, Param{}) {
		t.Fatalf("error when fetching single 1: %v", err)
	}

	t.Logf("received: %v", retSingle1)

	if !reflect.DeepEqual(retSingle1, cmp1) {
		t.Fatalf("received not equal to expected #1. Received: %v, expected: %v", retSingle1, cmp1)
	}

	var retSingle2 Param

	err = DB.Where("id = ?", fetchId).Take(&retSingle2).Error
	if err != nil || reflect.DeepEqual(retSingle2, Param{}) {
		t.Fatalf("error when fetching single 2: %v", err)
	}

	t.Logf("received: %v", retSingle2)

	if !reflect.DeepEqual(retSingle2, cmp1) {
		t.Fatalf("received not equal to expected #2. Received: %v, expected: %v", retSingle2, cmp1)
	}

	if !reflect.DeepEqual(retSingle2, retSingle1) {
		t.Fatalf("received not equal to expected #2. Received: %v, expected: %v", retSingle2, retSingle1)
	}

	var retMultiple []Param

	err = DB.Order("id ASC").Find(&retMultiple).Error
	if err != nil || reflect.DeepEqual(retMultiple, Param{}) {
		t.Fatalf("error when fetching multiple: %v", err)
	}

	t.Logf("received slice: %v", retMultiple)
	t.Logf("got previously #1: %v", retSingle1)
	t.Logf("got previously #2: %v", retSingle2)

	if len(retMultiple) != 2 {
		t.Fatalf("expected 2 items, got: %v", len(retMultiple))
	}

	if !reflect.DeepEqual(retSingle1, cmp1) {
		t.Fatalf("the received value is not equal to the expected value #1. Received: %v, expected: %v", retSingle1, cmp1)
	}

	if !reflect.DeepEqual(retSingle2, cmp1) {
		t.Fatalf("the received value is not equal to the expected value #2. Received: %v, expected: %v", retSingle2, cmp1)
	}

	if !reflect.DeepEqual(retMultiple[len(retMultiple)-1], cmp1) {
		t.Fatalf("the received slice item is not equal to the expected item. Received: %v, expected: %v", retMultiple[len(retMultiple)-1], cmp1)
	}
}
