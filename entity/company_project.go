package entity

import (
	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
	"time"
)

// QslModelUUID 全局模型公共参数 - 主键UUID模式
type QslModelUUID struct {
	ID             string         `gorm:"size:36;not null;primaryKey;comment:主键UUID"` // 主键UUID
	CreatedAt      time.Time      `gorm:"comment:创建时间"`
	CreatedAccount string         `gorm:"size:20;not null;default:'';comment:创建人账号"` // 创建人账号
	UpdatedAt      time.Time      `gorm:"comment:修改时间"`
	UpdatedAccount string         `gorm:"size:20;not null;default:'';comment:修改人账号"` // 修改人账号
	DeletedAt      gorm.DeletedAt `gorm:"index;comment:删除时间"`
	DeletedAccount string         `gorm:"size:20;not null;default:'';comment:删除人账号"` // 删除人账号
}

// CompanyProject 公司项目
type CompanyProject struct {
	QslModelUUID
	BindId       *string `gorm:"size:36;not null;default:'';comment:绑定ID"`
	Name         *string `gorm:"size:120;not null;default:'';comment:项目名称"`
	Department   *string `gorm:"size:50;not null;default:'';comment:所属部门"`
	Company      *string `gorm:"size:50;not null;default:'';comment:所属公司"`
	Description  *string `gorm:"size:500;not null;default:'';comment:项目描述"`
	ProjectAnnex *string `gorm:"size:90000;not null;default:'[]';comment:项目附件"`
	Content      *string `gorm:"size:90000;not null;default:'';comment:项目内容"`
	Sort         *uint   `gorm:"not null;default:0;comment:排序"`
}

func (*CompanyProject) TableName() string {
	return "company_project"
}

// BeforeUpdate 钩子 - 修改前自动填充修改人账号
func (c *CompanyProject) BeforeUpdate(tx *gorm.DB) (err error) {
	// 从 GORM 的 Context 中获取账号信息
	if ctxValue := tx.Statement.Context.Value("account"); ctxValue != nil {
		if account, ok := ctxValue.(string); ok {
			c.UpdatedAccount = account
		}
	}
	return
}

// BeforeCreate 钩子 - 创建前自动填充uuid和创建人账号
func (c *CompanyProject) BeforeCreate(tx *gorm.DB) (err error) {
	// 生成uuid模式的id
	if c.ID == "" {
		c.ID = uuid.Must(uuid.NewV4()).String()
	}
	// 生成uuid模式的bindId
	if c.BindId == nil {
		bindId := uuid.Must(uuid.NewV4()).String()
		c.BindId = &bindId
	}
	// 从 GORM 的 Context 中获取账号信息
	if ctxValue := tx.Statement.Context.Value("account"); ctxValue != nil {
		if account, ok := ctxValue.(string); ok {
			c.CreatedAccount = account
			c.UpdatedAccount = account
		}
	}
	return
}

// BeforeDelete 钩子 - 删除前自动填充删除人账号
func (c *CompanyProject) BeforeDelete(tx *gorm.DB) (err error) {
	// 从 GORM 的 Context 中获取账号信息
	if ctxValue := tx.Statement.Context.Value("account"); ctxValue != nil {
		if account, ok := ctxValue.(string); ok {
			c.DeletedAccount = account
		}
	}
	return
}
