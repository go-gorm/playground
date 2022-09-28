package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	testuser := []User{
		{
			ID:       "ID1",
			Revision: 1,
			Name:     "otherName",
			Embed: &Embed{
				FieldName: "Field",
			},
		},
		{
			ID:       "ID1",
			Revision: 2,
			Name:     "jinzhu",
			Embed: &Embed{
				FieldName: "Field",
			},
		},
		{
			ID:       "ID2",
			Revision: 3,
		},
		{
			ID:       "ID3",
			Revision: 4,
		},
		{
			ID:       "ID3",
			Revision: 5,
		},
	}
	for _, user := range testuser {
		err := DB.Create(&user).Error
		require.NoError(t, err)
	}

	users, err := getUsers([]string{"ID1", "ID3"})
	require.NoError(t, err)
	require.Len(t, users, 2)
	if users[0].ID == testuser[0].ID {
		assert.Equal(t, testuser[1], users[0])
		assert.Equal(t, testuser[4], users[1])
	} else {
		assert.Equal(t, testuser[1], users[0])
		assert.Equal(t, testuser[4], users[1])
	}
}

func getUsers(userIDs []string) ([]User, error) {
	var users []User
	subquery := DB.
		Select("MAX(revision) AS latest, id").
		Group("id").
		Where("users.id IN ?", userIDs).
		Table("users")
	err := DB.
		Table("users").
		Joins("RIGHT JOIN (?) AS r ON r.latest = users.revision", subquery).
		Find(&users).
		Error
	return users, err
}
