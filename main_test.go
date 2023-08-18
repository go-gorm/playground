package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite

func TestGORM(t *testing.T) {
	account := Account{
		ID: "account1",
		Network: Network{
			ID: "network1",
		},
		Peers: []Peer{
			{
				ID: "peer1",
			},
		},
	}

	DB.Session(&gorm.Session{FullSaveAssociations: true}).Create(account)

	var peer Peer
	if err := DB.First(&peer, "id = ?", account.Peers[0].ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	if peer.AccountID != account.ID {
		t.Error("account id of the peer doesn't match account id")
	}
}
