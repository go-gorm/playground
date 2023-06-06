package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	type TableStruct1 struct {
		ID   int
		Name string `gorm:"size:20;uniqueIndex"`
	}
	type TableStruct2 struct {
		ID   int
		Name string `gorm:"size:20;unique"`
	}
	type TableStruct3 struct {
		ID   int
		Name string `gorm:"size:20;index;unique"`
	}

	err := DB.Migrator().DropTable(&TableStruct1{}, &TableStruct2{}, &TableStruct3{})
	require.Empty(t, err)

	for i := 0; i < 5; i++ {
		err = DB.AutoMigrate(&TableStruct1{}, &TableStruct2{}, &TableStruct3{})
		require.Empty(t, err)
	}

	names1, err := GetIndexNames(&TableStruct1{})
	require.Empty(t, err)
	names2, err := GetIndexNames(&TableStruct2{})
	require.Empty(t, err)
	names3, err := GetIndexNames(&TableStruct3{})

	switch DB.Dialector.Name() {
	case "sqlite":
		assert.ElementsMatch(t, []string{"idx_table_struct1_name"}, names1)
		assert.ElementsMatch(t, []string{"sqlite_autoindex_table_struct2_1"}, names2)
		assert.ElementsMatch(t, []string{"idx_table_struct3_name", "sqlite_autoindex_table_struct3_1"}, names3)
	case "mysql":
		assert.ElementsMatch(t, []string{"PRIMARY", "idx_table_struct1_name"}, names1)
		assert.ElementsMatch(t, []string{"PRIMARY", "name"}, names2)
		assert.ElementsMatch(t, []string{"PRIMARY", "idx_table_struct3_name", "name"}, names3)
	case "postgres":
		assert.ElementsMatch(t, []string{"table_struct1_pkey", "idx_table_struct1_name"}, names1)
		assert.ElementsMatch(t, []string{"table_struct2_pkey", "table_struct2_name_key"}, names2)
		assert.ElementsMatch(t, []string{"table_struct3_pkey", "idx_table_struct3_name", "table_struct3_name_key"}, names3)
	}
}

func GetIndexNames(dst interface{}) ([]string, error) {
	if DB.Dialector.Name() == "sqlite" {
		stmt := &gorm.Statement{DB: DB}
		err := stmt.Parse(dst)
		if err != nil {
			return nil, err
		}

		var names []string
		err = DB.Raw("SELECT name FROM sqlite_master WHERE type = ? AND tbl_name = ?", "index", stmt.Table).Scan(&names).Error
		if err != nil {
			return nil, err
		}
		return names, nil
	} else {
		indexes, err := DB.Migrator().GetIndexes(dst)
		if err != nil {
			return nil, err
		}
		names := make([]string, 0, len(indexes))
		for _, index := range indexes {
			names = append(names, index.Name())
		}
		return names, nil
	}
}
