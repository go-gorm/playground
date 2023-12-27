package main

import (
	"reflect"
	"testing"
)

// GORM_REPO: ./gorm
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: `"jin'zhu"`} // string values contain single or double quotes

	// SQLite:
	// INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`name`,`age`,`birthday`,`company_id`,`manager_id`,`active`)
	// VALUES ("2023-12-27 17:58:17.329","2023-12-27 17:58:17.329",NULL,
	//   """jin'zhu""",
	//    0,NULL,NULL,NULL,false
	// ) RETURNING `id`
	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}

func TestExplain(t *testing.T) {
	type args struct {
		prepareSql string
		values     []interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantSQL string
	}{
		{"mysql", args{"SELECT ? AS QUOTES_STR", []interface{}{"'"}}, `SELECT '''' AS QUOTES_STR`},
		{"postgres", args{"SELECT $1 AS QUOTES_STR", []interface{}{"'"}}, `SELECT '''' AS QUOTES_STR`},
		{"sqlserver", args{"SELECT @p1 AS QUOTES_STR", []interface{}{"'"}}, `SELECT '''' AS QUOTES_STR`},
		{"sqlite", args{"SELECT ? AS QUOTES_STR", []interface{}{`"`}}, `SELECT """" AS QUOTES_STR`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if name := DB.Dialector.Name(); name != tt.name {
				t.Logf("%s skip %s...", name, tt.name)
				return
			}
			gotSQL := DB.Dialector.Explain(tt.args.prepareSql, tt.args.values...)
			if reflect.DeepEqual(gotSQL, tt.wantSQL) {
				var result string
				if err := DB.Raw(gotSQL).Row().Scan(&result); err == nil {
					t.Logf("exec `%s` result = `%s`", gotSQL, result)
				} else {
					t.Errorf("exec `%s` got error: %v", gotSQL, err)
				}
			} else {
				t.Errorf("Explain gotSQL = %v, want %v", gotSQL, tt.wantSQL)
			}
		})
	}
}
