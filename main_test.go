package main

import (
	"context"
	"testing"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite

type TenantOrder struct {
	ID         int64 `gorm:"primaryKey"`
	CustomerID int64
	Customer   TenantCustomer `gorm:"foreignKey:CustomerID;references:ID"`
}

func (TenantOrder) TableName() string {
	return "orders"
}

type TenantCustomer struct {
	ID   int64 `gorm:"primaryKey"`
	Name string
}

func (TenantCustomer) TableName() string {
	return "customers"
}

func TestGenericsAssociationJoinWithSchemaQualifiedTable(t *testing.T) {
	ctx := context.Background()

	mustExec(t, `ATTACH DATABASE ':memory:' AS tenant_1`)
	mustExec(t, `DROP TABLE IF EXISTS customers`)
	mustExec(t, `DROP TABLE IF EXISTS tenant_1.customers`)
	mustExec(t, `DROP TABLE IF EXISTS tenant_1.orders`)
	mustExec(t, `CREATE TABLE customers (id integer PRIMARY KEY, name text NOT NULL)`)
	mustExec(t, `CREATE TABLE tenant_1.customers (id integer PRIMARY KEY, name text NOT NULL)`)
	mustExec(t, `CREATE TABLE tenant_1.orders (id integer PRIMARY KEY, customer_id integer NOT NULL)`)
	mustExec(t, `INSERT INTO customers (id, name) VALUES (1, 'public customer')`)
	mustExec(t, `INSERT INTO tenant_1.customers (id, name) VALUES (1, 'tenant customer')`)
	mustExec(t, `INSERT INTO tenant_1.orders (id, customer_id) VALUES (1, 1)`)

	got, err := gorm.G[TenantOrder](DB).
		Table("tenant_1.orders").
		Joins(clause.InnerJoin.Association("Customer"), nil).
		First(ctx)
	if err != nil {
		t.Fatalf("association join failed: %v", err)
	}
	if got.Customer.Name != "tenant customer" {
		t.Errorf("association join used the wrong physical table, got customer %q, want %q", got.Customer.Name, "tenant customer")
	}

	workaround, err := gorm.G[TenantOrder](DB).
		Table("tenant_1.orders").
		Joins(
			clause.InnerJoin.AssociationFrom(
				"Customer",
				clause.Expr{SQL: "SELECT * FROM tenant_1.customers"},
			).As("Customer"),
			nil,
		).
		First(ctx)
	if err != nil {
		t.Fatalf("association join workaround failed: %v", err)
	}
	if workaround.Customer.Name != "tenant customer" {
		t.Fatalf("association join workaround got customer %q, want %q", workaround.Customer.Name, "tenant customer")
	}
}

func mustExec(t *testing.T, query string) {
	t.Helper()

	if err := DB.Exec(query).Error; err != nil {
		t.Fatalf("exec %q: %v", query, err)
	}
}
