package main

import (
	"strings"

	"gorm.io/gorm/schema"
)

type NamingStrategy struct {
	baseStrategy schema.NamingStrategy
}

func (ns NamingStrategy) TableName(table string) string {
	return ns.baseStrategy.TableName(table)
}

func (ns NamingStrategy) ColumnName(table, column string) string {
	baseColumnName := ns.baseStrategy.ColumnName(table, column)

	if table == "" {
		return baseColumnName
	}

	s := strings.Split(table, "_")

	var prefix string
	switch len(s) {
	case 1:
		prefix = s[0][:3]
	case 2:
		prefix = s[0][:1] + s[1][:2]
	default:
		prefix = s[0][:1] + s[1][:1] + s[2][:1]
	}
	return prefix + "_" + baseColumnName
}

func (ns NamingStrategy) JoinTableName(table string) string {
	return ns.baseStrategy.JoinTableName(table)
}

func (ns NamingStrategy) RelationshipFKName(rel schema.Relationship) string {
	return ns.baseStrategy.RelationshipFKName(rel)
}

func (ns NamingStrategy) CheckerName(table, column string) string {
	return ns.baseStrategy.CheckerName(table, column)
}

func (ns NamingStrategy) IndexName(table, column string) string {
	return ns.baseStrategy.IndexName(table, column)
}
