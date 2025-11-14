package main

import "fmt"

type Tenant struct {
	gorm.Model

	Description string
}

type BaseModel struct {
	gorm.Model

	TenantID uint `gorm:"primaryKey"`
	Tenant   Tenant
}

type Player struct {
	BaseModel

	Name string
	Teams []Team
}

type Team struct {
	BaseModel

	Name     string
	PlayerID uint
}

func main() {
	DB.Automigrate(
		&Tenant{},
		&Player{},
		&Team{},
	)
}
