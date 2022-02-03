package main

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: postgres

type Event struct {
	gorm.Model
	ID  string `gorm:"primaryKey"`
	UID uint32 `gorm:"not null;autoIncrement"`
}

func TestGORM(t *testing.T) {
	require.NoError(t, DB.Migrator().DropTable(&Event{}))
	require.NoError(t, DB.AutoMigrate(&Event{}))
	require.NoError(t, DB.AutoMigrate(&Event{}))

	require.NoError(t, DB.Save(&Event{ID: "a"}).Error)
	require.NoError(t, DB.Save(&Event{ID: "b"}).Error)

	events := make([]*Event, 0, 2)
	DB.Find(&events)

	require.Equal(t, 2, len(events))
	require.NotEqual(t, events[0].UID, events[1].UID)
}
