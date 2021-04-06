package main

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	AuditedModel
	Name        string
	DOB         time.Time
	CreditCards []CreditCard
}

type CreditCard struct {
	gorm.Model
	AuditedModel
	Number string
	UserID uint
}