package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	accountSender := &Account{Email: "test@test.test"}
	if err := DB.Save(accountSender).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	customerSender := &Customer{Name: "test"}
	if err := DB.Save(customerSender).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	accountMessage := &TicketMessage{
		Message:       "hello world!",
		AccountSender: accountSender,
	}
	if err := DB.Save(accountMessage).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	customerMessage := &TicketMessage{
		Message:        "hello world!",
		CustomerSender: customerSender,
	}
	if err := DB.Save(customerMessage).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
