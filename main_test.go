package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	// Migrate the schema
	DB.AutoMigrate(&Product{})
	DB.AutoMigrate(&ProductNG3{})

	
	var isImput int

	isImput = 0 // v2 版本也存在此问题  不需要归集   巨坑: 字段默认值 IsImput 为 -1 (负数 -2 )  插入 0  插入的值仍然为 -1 (-2)

	// Create
	DB.Create(&Product{Code: "L1212", Price: 1000, IsImput: isImput})
	DB.Create(&ProductNG3{Code: "L1213", Price: 1000, IsImput: -1})

	var product Product
	DB.First(&product, "code = ?", "L1212") // find product with code l1212
	if product.IsImput != 0 {
		t.Errorf("product IsImput is %v, expect : 0", product.IsImput)
	}

	var productng3 ProductNG3
	DB.First(&productng3, "code = ?", "L1213") // find product with code l1213
	if productng3.IsImput != -1 {
		t.Errorf("product IsImput is %v, expect : -1", product.IsImput)
	}

}
