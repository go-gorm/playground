package main

import (
	"fmt"
	"testing"

	"gorm.io/gorm"
	"gorm.io/playground/models"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := models.User{Name: "jinzhu"}

	DB.Create(&user)

	var result models.User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}

// TestBug_SessionPrepareStmtCannotDowngrade asserts that
// db.Session(&gorm.Session{PrepareStmt: false}) on a parent opened with
// PrepareStmt:true should yield a child whose Statement.ConnPool is no
// longer a *gorm.PreparedStmtDB / *gorm.PreparedStmtTX.
//
// Today the assertion FAILS — the explicit `false` is silently dropped in
// gorm.go::Session() because there is no `else` branch to undo the
// inherited prepared pool. The bool zero-value cannot be distinguished
// from "caller wants false".
//
// Source path: https://github.com/go-gorm/gorm/blob/v1.31.1/gorm.go#L280-L314
//
// Prior closed-without-fix issues describing the same defect:
//   - https://github.com/go-gorm/gorm/issues/3362  (closed 2020-09 as stale)
//   - https://github.com/go-gorm/gorm/issues/4991  (closed 2022-01 as feature_request)
//   - https://github.com/go-gorm/gorm/issues/5026  (closed 2022-01 as question)
//
// This test does NOT depend on the postgres driver — it is observable
// purely from Statement.ConnPool type inspection on any dialect.
func TestBug_SessionPrepareStmtCannotDowngrade(t *testing.T) {
	// Force a parent with PrepareStmt: true. We can't reopen the playground's
	// shared DB, so build a sibling session that wraps the existing pool with
	// a PreparedStmtDB — equivalent to having opened with PrepareStmt:true.
	parent := DB.Session(&gorm.Session{PrepareStmt: true})

	parentPoolType := fmt.Sprintf("%T", parent.Statement.ConnPool)
	if parentPoolType != "*gorm.PreparedStmtDB" && parentPoolType != "*gorm.PreparedStmtTX" {
		t.Fatalf("setup precondition failed: expected parent ConnPool to be *PreparedStmtDB/TX, got %s", parentPoolType)
	}

	// The user request: downgrade for a child session.
	child := parent.Session(&gorm.Session{PrepareStmt: false})

	childPoolType := fmt.Sprintf("%T", child.Statement.ConnPool)
	t.Logf("parent ConnPool: %s", parentPoolType)
	t.Logf("child  ConnPool: %s", childPoolType)

	if childPoolType == "*gorm.PreparedStmtDB" || childPoolType == "*gorm.PreparedStmtTX" {
		t.Fatalf(
			"BUG REPRODUCED: child Session(&gorm.Session{PrepareStmt:false}) still uses %s — the explicit false was silently dropped",
			childPoolType,
		)
	}
}
