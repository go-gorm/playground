package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	l := Layer{
			Labels: []*Label{
				{
					Name: "label 1",
				},
				{
					Name: "label 2",
				},
			},
	}
	t.Log("CREATE MAP")
	err := DB.Create(&l).Error
	require.NoError(t, err)

	ps := []*Point{
		{
			Layer:   &l,
			Labels: []*Label{
				l.Labels[0],
				l.Labels[1],
			},
		},
		{
			Layer:   &l,
			Labels: []*Label{
				l.Labels[1],
			},
		},
	}

	t.Log("CREATE POINTS")
	err = DB.Create(&ps).Error
	require.NoError(t, err)
}
