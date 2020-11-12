package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	type Hoge struct {
		ID   uint
		Name string
	}
	type Fuga struct {
		ID   uint
		Name string
	}

	if err := DB.AutoMigrate(&Hoge{}, &Fuga{}); err != nil {
		t.Errorf("%v", err)
	}

	DB = DB.Session(&gorm.Session{})
	DB = DB.InstanceSet("key", "value") // Correct without this line
	// DB = DB.Debug() // Correct with this line
	if err := DB.Create(&Hoge{Name: "Alice"}).Error; err != nil {
		t.Errorf("%v", err)
	}
	if err := DB.Create(&Fuga{Name: "Bob"}).Error; err != nil {
		t.Errorf("%v", err)
	}
	if err := DB.Where("1=1").Delete(&Fuga{}).Error; err != nil {
		t.Errorf("%v", err)
	}
	if err := DB.Where("1=1").Delete(&Hoge{}).Error; err != nil {
		t.Errorf("%v", err)
	}

	// 2020/11/12 17:17:28 /Users/soranoba/go/pkg/mod/gorm.io/driver/sqlite@v1.1.3/migrator.go:32
	// [0.031ms] [rows:-] SELECT count(*) FROM sqlite_master WHERE type='table' AND name="hoges"

	// 2020/11/12 17:17:28 /Users/soranoba/Documents/playground/main_test.go:23
	// [0.185ms] [rows:-] SELECT * FROM `hoges` LIMIT 1

	// 2020/11/12 17:17:28 /Users/soranoba/go/pkg/mod/gorm.io/driver/sqlite@v1.1.3/migrator.go:32
	// [0.029ms] [rows:-] SELECT count(*) FROM sqlite_master WHERE type='table' AND name="fugas"

	// 2020/11/12 17:17:28 /Users/soranoba/Documents/playground/main_test.go:23
	// [0.068ms] [rows:-] SELECT * FROM `fugas` LIMIT 1

	// 2020/11/12 17:17:28 /Users/soranoba/Documents/playground/main_test.go:30
	// [1.384ms] [rows:1] INSERT INTO `hoges` (`name`) VALUES ("Alice")

	// 2020/11/12 17:17:28 /Users/soranoba/Documents/playground/main_test.go:33
	// [1.088ms] [rows:1] INSERT INTO `hoges` (`name`) VALUES ("Bob")

	// 2020/11/12 17:17:28 /Users/soranoba/Documents/playground/main_test.go:36
	// [0.990ms] [rows:2] DELETE FROM `hoges` WHERE 1=1

	// 2020/11/12 17:17:28 /Users/soranoba/Documents/playground/main_test.go:39
	// [0.351ms] [rows:0] DELETE FROM `hoges` WHERE 1=1 AND 1=1

	//           !?!?!?!?!?!?!
}
