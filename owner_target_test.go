package main

import (
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Target 目标模型：包含主键 ID 与一个同名的非主键列 FKID（映射到列 fk_id）
// 注意：这里的 FKID 不是主键，仅用于模拟“目标表中存在与外键同名的非主键列”的场景
type Target struct {
	ID   int  `gorm:"primaryKey"`
	FKID *int `gorm:"column:fk_id"`
}

func (Target) TableName() string { return "targets" }

// Owner 所有者模型：通过外键 FKID（映射到列 fk_id）belongs_to Target
// 显式指定 foreignKey 与 references，期望生成关联条件：owners.fk_id = targets.id
// 回归现象：被误判为 has one，关联条件变为 owners.fk_id = targets.fk_id（错误）
type Owner struct {
	ID     int     `gorm:"primaryKey"`
	FKID   int     `gorm:"column:fk_id"`
	Target *Target `gorm:"foreignKey:FKID;references:ID"`
}

func (Owner) TableName() string { return "owners" }

// TestBelongsToSameNamedColumnRegression 验证 belongs_to 在“目标模型存在同名非主键列”时的回归问题
// 用最小数据集复现：
// 1) 目标表插入一条记录：ID=1，FKID=NULL（非主键）
// 2) 所有者表插入一条记录：FKID=1（外键指向目标主键）
// 3) 预加载 Owner.Target，预期应拿到 Target.ID==1
// 若发生回归，JOIN/ON 条件将错误使用 targets.fk_id（非主键），导致取到 nil 或错误记录
func TestBelongsToSameNamedColumnRegression(t *testing.T) {
	// 使用 sqlite 内存数据库，避免外部依赖
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("open db err: %v", err)
	}

	// 自动迁移模型结构
	if err := db.AutoMigrate(&Target{}, &Owner{}); err != nil {
		t.Fatalf("migrate err: %v", err)
	}

	// 插入目标记录：主键 ID=1；非主键同名列 FKID 为空
	if err := db.Create(&Target{ID: 1, FKID: nil}).Error; err != nil {
		t.Fatalf("create target err: %v", err)
	}
	// 插入所有者记录：外键 FKID=1（按 belongs_to 应关联到 targets.id=1）
	if err := db.Create(&Owner{ID: 100, FKID: 1}).Error; err != nil {
		t.Fatalf("create owner err: %v", err)
	}

	// 预加载 belongs_to 关联
	var owners []Owner
	if err := db.Preload("Target").Find(&owners).Error; err != nil {
		t.Fatalf("query err: %v", err)
	}

	// 断言：应拿到 Target.ID==1
	// 如果被误判为 has one 并按 targets.fk_id 关联，可能得到 nil 或错误记录
	if len(owners) == 0 || owners[0].Target == nil || owners[0].Target.ID != 1 {
		t.Fatalf("expected target.id=1, got=%v", owners[0].Target)
	}
}
