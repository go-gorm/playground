package main

import (
	"gorm.io/gorm/clause"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS:  mysql

const (
	_CreatedTableSql = `CREATE TABLE location (
id int unsigned NOT NULL AUTO_INCREMENT,
ip_start int unsigned NOT NULL DEFAULT '0',
ip_end int unsigned NOT NULL DEFAULT '0',
ip_range linestring NOT NULL SRID 3857,
PRIMARY KEY (id),
SPATIAL KEY si_geo_ip_range (ip_range)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;`

	_CreatedRowSql = "INSERT INTO location (ip_start,ip_end,ip_range) VALUES (8060929,8060930,ST_GeomFromText('LINESTRING(8060929 -1, 8060930 1)', 3857))"
)

func TestGORM(t *testing.T) {

	// err = nil
	if err := DB.Exec(_CreatedTableSql).Error; err != nil {
		t.Fatalf("created table fail, err:%v", err)
	}

	// err = nil
	if err := DB.Exec(_CreatedRowSql).Error; err != nil {
		t.Fatalf("created location fail, err:%v", err)
	}

	err := DB.Table("location").Create(map[string]interface{}{
		"ip_start": 8060929,
		"ip_end":   8060930,
		"ip_range": clause.Expr{
			SQL:  "ST_GeomFromText('LINESTRING(? -1, ? 1)', 3857)",
			Vars: []interface{}{8060929, 8060930},
		},
	}).Error

	if err != nil {
		t.Errorf("user API created location fail, err:%v", err)
	}
}
