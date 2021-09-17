package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	type Context struct {
		CompanyID *int
		ManagerID *uint
	}

	var (
		contextKey string = "APP-CONTEXT"
		companyID  int    = 1
		managerID  uint   = 1
	)

	assignContext := func(db *gorm.DB) {
		if ICtx, ok := db.Get(contextKey); ok {
			ctx := ICtx.(*Context)

			if db.Statement.Schema != nil {

				if field := db.Statement.Schema.LookUpField("CompanyID"); field != nil {
					if v, isZero := field.ValueOf(db.Statement.ReflectValue); v == nil || isZero {
						field.Set(db.Statement.ReflectValue, ctx.CompanyID)
					}
				}

				if field := db.Statement.Schema.LookUpField("ManagerID"); field != nil {
					if v, isZero := field.ValueOf(db.Statement.ReflectValue); v == nil || isZero {
						field.Set(db.Statement.ReflectValue, ctx.CompanyID)
					}
				}

			}
		}
	}

	callback := DB.Callback()
	callback.Create().After("gorm:before_create").Register("gobp:assign_tenant", assignContext)

	user := User{
		Name: "User",
	}

	ctx := Context{CompanyID: &companyID, ManagerID: &managerID}

	DB.Set(contextKey, &ctx).Create(&user)

	var result User
	if err := DB.Where(&User{CompanyID: &companyID, ManagerID: &managerID}).First(&result).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
