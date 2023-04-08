package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type Index struct {
	Table 		string
	Key_name	string
	Non_unique	uint
	Column_name	string
	Seq_in_index	uint
}

func TestGORM(t *testing.T) {
	t.Logf("Running AutoMigrate multiple times...\n");
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&User{})
	rows, err := DB.Raw("SHOW INDEX FROM users").Rows()
	defer rows.Close()
	if err != nil {
		t.Fatalf("Failed on query, got error: %v", err)
	}
	var index Index
	count := 0
	for rows.Next() {
		err = DB.ScanRows(rows, &index)
		if err != nil {
			t.Fatalf("Failed on scan, got error: %v", err)
		}
		t.Logf("index: %+v\n", index)
		count += 1
	}
	if count != 1 {
		t.Fatalf("Created multiple unique indexes on the same field when executing AutoMigrate, count: %v\n", count)
	}
}
