package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: mysql
// __TEST_DRIVERS__: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	appInitial := &App{
		Name: "foo",
		Envs: []*Env{
			{Key: "KEY1", Value: "value1"},
		},
	}
	DB.Create(appInitial)
	t.Logf("appInitial: %+v", appInitial)
	for i, env := range appInitial.Envs {
		t.Logf("  env[%d]: %+v", i, env)
	}

	appUpdated := &App{
		ID:   appInitial.ID,
		Name: "foo",
		Envs: []*Env{
			{Key: "KEY1", Value: "value1-updated"},
			{Key: "KEY2", Value: "value2-added"},
		},
	}
	DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(appUpdated)
	t.Logf("appUpdated: %+v", appUpdated)
	for i, env := range appUpdated.Envs {
		t.Logf("  env[%d]: %+v", i, env)
	}

	appFetched := &App{ID: appInitial.ID}
	DB.Preload("Envs").First(appFetched)
	t.Logf("appFetched: %+v", appFetched)
	for i, env := range appFetched.Envs {
		t.Logf("  env[%d]: %+v", i, env)
	}

	for i := range appFetched.Envs {
		envFetched := appFetched.Envs[i]
		envUpdated := appUpdated.Envs[i]

		if envFetched.Key != envUpdated.Key {
			t.Errorf("Expected env key %s, got %s", envFetched.Key, envUpdated.Key)
		}
		if envFetched.Value != envUpdated.Value {
			t.Errorf("Expected env value %s, got %s", envFetched.Value, envUpdated.Value)
		}
		if envFetched.ID != envUpdated.ID {
			t.Errorf("Expected env ID %d, got %d", envFetched.ID, envUpdated.ID)
		}
	}
}
