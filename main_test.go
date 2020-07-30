package main

import (
	"testing"
	"github.com/gofrs/uuid"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type HostGroup struct {
	ID uuid.UUID `json:"id,omitempty" gorm:"type:uuid;primary_key;"`

	Hosts []*Host `json:"omitempty" gorm:"foreignkey:HostGroupID; constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
}

type Host struct {
	ID uuid.UUID `json:"id,omitempty" gorm:"type:uuid;primary_key;"`

	// The ID of a host group that the host belongs to.
	HostGroupID *uuid.UUID `json:"host_group_id,omitempty" gorm:"type:uuid"`
}



func TestGORM(t *testing.T) {
	DB.AutoMigrate(&HostGroup{})
	DB.AutoMigrate(&Host{})
	DB.Exec("INSERT INTO hosts (id) VALUES ('22222222-2222-2222-2222-222222222222')") //Foreign key null
	hosts := make([]*host.Host, 0)
	err := DB.Find(&hosts).Error
	if err != nil {
		panic(err)
	}
}
