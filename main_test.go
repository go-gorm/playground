package main

import (
	"fmt"
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: postgres

const testEntName = "test_pent"

type PrimitiveEntity struct {
	gorm.Model
}

func (pe PrimitiveEntity) TableName() string {
	return testEntName
}

func TestGORM(t *testing.T) {
	if DB.Dialector.Name() != "postgres" {
		t.Fatalf("This is postgres-only specific test, not tested on other databases")
	}

	DB.Migrator().DropTable(&PrimitiveEntity{})

	if err := DB.AutoMigrate(&PrimitiveEntity{}); err != nil {
		t.Fatalf("Failed to migrate for PrimitiveEntity entity: %v", err)
	}

	pEnt := PrimitiveEntity{}
	DB.Create(&pEnt)

	var checkUser PrimitiveEntity
	if err := DB.First(&checkUser, pEnt.ID).Error; err != nil {
		t.Errorf("Failed to seek for User entity, got error: %v", err)
	}

	// Primary key info
	type PKInfo struct {
		AttName    string `gorm:"column:attname"`
		FormatType string `gorm:"column:format_type"`
	}

	var pePK PKInfo
	tx := DB.Raw(
		fmt.Sprintf(`
		SELECT
			a.attname,
			format_type(a.atttypid, a.atttypmod) 
		FROM
			pg_attribute a
			JOIN (SELECT *, GENERATE_SUBSCRIPTS(indkey, 1) AS indkey_subscript FROM pg_index) AS i
			ON
				i.indisprimary
				AND i.indrelid = a.attrelid
				AND a.attnum = i.indkey[i.indkey_subscript]
		WHERE
			a.attrelid = '%s'::regclass
		ORDER BY
			i.indkey_subscript;`, testEntName),
		2).Scan(&pePK)

	if tx.Error != nil {
		t.Fatalf("Failed to get primary key info for PrimitiveEntity entity: %v", tx.Error)
	}

	// Primary key should be present
	if pePK.AttName != "id" ||
		pePK.FormatType != "bigint" {
		t.Fatalf("PrimitiveEntity primary key unexpected format: %s, %s",
			pePK.AttName,
			pePK.FormatType)
	}

	// Index info
	type IndexInfo struct {
		TableName  string `gorm:"column:table_name"`
		IndexName  string `gorm:"column:index_name"`
		ColumnName string `gorm:"column:column_name"`
	}

	var indxInfo IndexInfo
	tx = DB.Raw(
		fmt.Sprintf(`
		SELECT
			t.relname AS table_name,
			i.relname AS index_name,
			a.attname AS column_name
		FROM
			pg_class t,
			pg_class i,
			pg_index ix,
			pg_attribute a
		WHERE
			t.oid = ix.indrelid
			and i.oid = ix.indexrelid
			and a.attrelid = t.oid
			and a.attnum = ANY(ix.indkey)
			and t.relkind = 'r'
			and t.relname = '%s'
			and a.attname = 'id'
		ORDER BY
			t.relname,
			i.relname;`, testEntName),
		3).Scan(&indxInfo)

	if tx.Error != nil {
		t.Fatalf("Failed to get index info for PrimitiveEntity, reason: %v", tx.Error)
	}

	// While Index should NOT be present
	if indxInfo.TableName != "" ||
		indxInfo.IndexName != "" ||
		indxInfo.ColumnName != "" {
		t.Fatalf("Unexpected PrimitiveEntity index presence: %s, %s, %s",
			indxInfo.TableName,
			indxInfo.IndexName,
			indxInfo.ColumnName)
	}
}
