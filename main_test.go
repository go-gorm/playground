package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	initDB(t)

	expectedBars := map[int]int{
		0: 2,
		1: 3,
	}

	for fooID, expectedBarCount := range expectedBars {
		var foo Foo
		tx := DB.
			Preload("Bars").
			Where("foo_id = ?", fooID).
			Find(&foo)

		if err := tx.Error; err != nil {
			t.Errorf("Failed, got error: %v", err)
		}

		println("- foo", foo.FooID, foo.Value)
		for _, bar := range foo.Bars {
			println("  * bar", bar.BarID, bar.Value)
		}

		if len(foo.Bars) != expectedBarCount {
			t.Errorf("Failed for FooID=%d, expected %d bars, but got %d", fooID, expectedBarCount, len(foo.Bars))
		}
	}

}

func initDB(t *testing.T) {
	dbExec(t, "DROP TABLE IF EXISTS foo_bar")
	dbExec(t, "DROP TABLE IF EXISTS foo")
	dbExec(t, "DROP TABLE IF EXISTS bar")

	dbExec(t, `
		CREATE TABLE foo (
			foo_id INTEGER NOT NULL,
			value TEXT NOT NULL,
			PRIMARY KEY (foo_id)
		)
	`)
	dbExec(t, `
		CREATE TABLE bar (
			bar_id INTEGER NOT NULL,
			value TEXT NOT NULL,
			PRIMARY KEY (bar_id)
		)
	`)
	dbExec(t, `
		CREATE TABLE foo_bar (
			foo_id INTEGER NOT NULL REFERENCES foo(foo_id),
			bar_id INTEGER NOT NULL REFERENCES bar(bar_id),
			PRIMARY KEY (foo_id, bar_id)
		)
	`)

	// foos
	dbExec(t, "INSERT INTO foo (foo_id, value) VALUES (0, 'foo0')")
	dbExec(t, "INSERT INTO foo (foo_id, value) VALUES (1, 'foo1')")

	// bars
	dbExec(t, "INSERT INTO bar (bar_id, value) VALUES (0, 'bar0')")
	dbExec(t, "INSERT INTO bar (bar_id, value) VALUES (1, 'bar1')")
	dbExec(t, "INSERT INTO bar (bar_id, value) VALUES (2, 'bar2')")

	// foo0 has 2 bars
	dbExec(t, "INSERT INTO foo_bar (foo_id, bar_id) VALUES (0, 0)")
	dbExec(t, "INSERT INTO foo_bar (foo_id, bar_id) VALUES (0, 1)")

	// foo1 has 3 bars
	dbExec(t, "INSERT INTO foo_bar (foo_id, bar_id) VALUES (1, 0)")
	dbExec(t, "INSERT INTO foo_bar (foo_id, bar_id) VALUES (1, 1)")
	dbExec(t, "INSERT INTO foo_bar (foo_id, bar_id) VALUES (1, 2)")
}

func dbExec(t *testing.T, sql string) {
	if err := DB.Exec(sql).Error; err != nil {
		t.Errorf("Failed exec sql, got error: %v", err)
	}
}

type Foo struct {
	FooID uint32 `gorm:"column:foo_id;autoIncrement:false"`
	Value string `gorm:"column:value"`

	Bars []*Bar `gorm:"many2many:foo_bar;joinReferences:bar_id;joinForeignKey:foo_id;references:bar_id;foreignKey:foo_id"`
}

func (Foo) TableName() string {
	return "foo"
}

type Bar struct {
	BarID uint32 `gorm:"column:bar_id;autoIncrement:false"`
	Value string `gorm:"column:value"`
}

func (Bar) TableName() string {
	return "bar"
}
