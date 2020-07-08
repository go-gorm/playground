package main

import (
	"testing"
)

type Action struct {
	IdAction int64  `gorm:"column:idaction;primary_key"`
	Roles    []Role `gorm:"Many2Many:roleaction;JoinForeignKey:idaction;JoinReferences:idrole"`
}

type Role struct {
	IdRole  int64    `gorm:"column:idrole;primary_key"`
	Actions []Action `gorm:"Many2Many:roleaction;JoinForeignKey:idrole;JoinReferences:idaction"`
}

type RoleAction struct {
	IdRole   int64 `gorm:"column:idrole;primary_key"`
	IdAction int64 `gorm:"column:idaction;primary_key"`
}

func TestGORM(t *testing.T) {

	var roles []Role

	if err := DB.AutoMigrate(Action{}, Role{}, RoleAction{}); err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if err := DB.Model(Action{}).Association("Roles").Find(&roles); err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

}
