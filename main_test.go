package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres

func TestGORM(t *testing.T) {

	givenUser(t, User2{ID: 1, Name: "user1"})
	givenProfile(t, Profile{ID: 1, UserID: 1, Description: "profile1"})

	var p Profile
	err := DB.Debug().Preload("User").Find(&p).Error
	require.NoError(t, err)
	require.Equal(t, "user1", p.User.Name)
}

func givenProfile(t *testing.T, profile Profile) {
	err := DB.Create(&profile).Error
	require.NoError(t, err)
}

func givenUser(t *testing.T, user User2) {
	err := DB.Create(&user).Error
	require.NoError(t, err)
}
