package main

import "fmt"

const (
	DbTablePrefix = "_wt_"
)

type Table struct {
	name string
}

// Init 表格名称刷新
func (t Table) Init(name string) Table {
	t.name = GetTableNaming(name)
	return t
}

// Naming 获取表明
func (t Table) Naming() string {
	return t.name
}

// Field 获取查询字段
func (t Table) Field(fields ...string) []string {
	var fld []string
	for _, f := range fields {
		fld = append(fld, fmt.Sprintf("%v.%v", t.name, f))
	}
	return fld
}


func GetTableNaming(tb string) string {
	return fmt.Sprintf("%v%v", DbTablePrefix, tb)
}
