package main

import (
	"strconv"
	"testing"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}

type Detail struct {
	gorm.Model
	OrgID      string `gorm:"index;comment:组织id"`
	OrderID    string `gorm:"unique;index;uniqueIndex;comment:订单id"`
	PID        string `gorm:"index;comment:产品id"`
	Apply      int    `gorm:"comment:申请"`
	CreditSucc int    `gorm:"comment:审批"`
	CreditFail int    `gorm:"comment:审批"`
	Curr       int    `gorm:"comment:授信"`
	Extract    int    `gorm:"comment:提款"`
}

func ds() int {
	t := time.Now()
	return int(time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix())
}

func (d *Detail) AfterCreate(tx *gorm.DB) error {
	t := ds()
	s := Stat{
		Time:       t,
		OrgID:      d.OrgID,
		PID:        d.PID,
		Work:       d.CreditSucc + d.CreditFail,
		Apply:      d.Apply,
		CreditSucc: d.CreditSucc,
		CreditFail: d.CreditFail,
		Curr:       d.Curr,
		Extract:    d.Extract,
	}
	sm := map[string]interface{}{
		"work":        gorm.Expr("work + ?", s.Work),
		"apply":       gorm.Expr("apply+?", s.Apply),
		"credit_succ": gorm.Expr("credit_succ+?", s.CreditSucc),
		"credit_fail": gorm.Expr("credit_fail + ?", s.CreditFail),
		"curr":        gorm.Expr("curr + ?", s.Curr),
		"extract":     gorm.Expr("extract + ?", s.Extract),
	}

	if err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "time"}, {Name: "org_id"}, {Name: "p_id"}},
		DoUpdates: clause.Assignments(sm),
	}).Create(&s).Error; err != nil {
		return err
	}

	return nil
}

type Stat struct {
	gorm.Model
	Time       int    `gorm:"index;comment:时间" pg:"unique:s"`
	OrgID      string `gorm:"index;comment:组织id" pg:"unique:s"`
	PID        string `gorm:"index;comment:产品id" pg:"unique:s"`
	Apply      int    `gorm:"comment:申请"`
	Work       int    `gorm:"comment:处理"`
	CreditSucc int    `gorm:"comment:审批"`
	CreditFail int    `gorm:"comment:审批"`
	Curr       int    `gorm:"comment:授信"`
	Extract    int    `gorm:"comment:提款"`
}

func TestGORM2(t *testing.T) {
	if err := DB.AutoMigrate(new(Detail), new(Stat)); err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	if !DB.Migrator().HasIndex(new(Stat), "stats_a") {
		if err := DB.Exec("CREATE UNIQUE INDEX stats_a on stats (time,org_id,p_id)").Error; err != nil {
			t.Errorf("Failed, got error: %v", err)
		}
	}
	if err := DB.Create(&Detail{
		OrgID:      "a",
		PID:        "b",
		OrderID:    strconv.Itoa(int(time.Now().UnixNano())),
		Apply:      1,
		CreditSucc: 0,
		CreditFail: 0,
		Curr:       0,
		Extract:    0,
	}).Error; err != nil {
		panic(err)
	}
}
