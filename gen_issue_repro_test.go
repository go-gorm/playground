package main

import (
	"strings"
	"testing"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

type Membership struct {
	GroupID string `gorm:"primaryKey"`
	UserID  string `gorm:"primaryKey"`
}

func (Membership) TableName() string { return "memberships" }

func TestUpdateLeaksPrimaryKeyIntoSharedModel(t *testing.T) {
	if DB.Dialector.Name() != "postgres" {
		t.Skip("this reproduction requires PostgreSQL")
	}

	if err := DB.Migrator().DropTable(&Membership{}); err != nil {
		t.Fatal(err)
	}
	if err := DB.AutoMigrate(&Membership{}); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := DB.Migrator().DropTable(&Membership{}); err != nil {
			t.Errorf("clean up memberships: %v", err)
		}
	})

	groupID := field.NewString("memberships", "group_id")
	userID := field.NewString("memberships", "user_id")

	model := &Membership{}
	memberships := new(gen.DO)
	memberships.UseDB(DB)
	memberships.UseModel(model)

	subQuery := func(group string) *gorm.Statement {
		sub := memberships.WithContext(t.Context()).Where(groupID.Eq(group)).Select(userID)
		return DB.Session(&gorm.Session{DryRun: true}).Table("users").
			Where("id IN (?)", sub.(*gen.DO).UnderlyingDB()).
			Find(&[]map[string]any{}).Statement
	}

	before := subQuery("A")
	if before.Error != nil {
		t.Fatal(before.Error)
	}
	beforeSQL := before.SQL.String()
	t.Logf("before Update: %s; vars: %v", beforeSQL, before.Vars)
	if strings.Count(beforeSQL, `"group_id" =`) != 1 {
		t.Fatalf("expected one group_id condition before Update: %s", beforeSQL)
	}

	// No rows are needed: the shared model must stay unchanged even for an empty update.
	if _, err := memberships.WithContext(t.Context()).Where(groupID.Eq("X")).Update(groupID, "Y"); err != nil {
		t.Fatal(err)
	}

	after := subQuery("B")
	if after.Error != nil {
		t.Fatal(after.Error)
	}
	afterSQL := after.SQL.String()
	t.Logf("after Update:  %s; vars: %v; shared model: %+v", afterSQL, after.Vars, model)
	if afterSQL != beforeSQL || strings.Count(afterSQL, `"group_id" =`) != 1 {
		t.Errorf("value from Update ('Y') leaked into later subquery: %s; vars: %v", afterSQL, after.Vars)
	}
	if len(after.Vars) != 1 || after.Vars[0] != "B" {
		t.Errorf("expected only the new group filter ('B'), got vars: %v", after.Vars)
	}
	if *model != (Membership{}) {
		t.Errorf("Update mutated the shared model registered via UseModel: %+v", model)
	}
}
