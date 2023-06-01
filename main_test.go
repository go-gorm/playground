package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	var dists []Distribution
	places := []Place{
		{IdentityKey: "foo"},
		{IdentityKey: "bar"},
	}
	DB.Create(&places[0])
	DB.Create(&places[1])
	DB.Create(Distribution{ID: []byte("some-id"), Places: places, IdentityKey:"baz"})

	err := DB.Debug().Model(&Distribution{}).
			Where("places.id IN (?)", []uint{places[0].ID, places[1].ID}).
			Association("Places").
			Find(&dists)
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	if len(dists) == 0 {
		t.Errorf("No distributions received")
	}
}
