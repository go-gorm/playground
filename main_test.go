package main

import (
	"testing"

	"github.com/hashicorp/go-secure-stdlib/base62"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	id, err := base62.Random(10)
	if err != nil {
		t.Fatalf("unable to generate id: %v", err)
	}
	id2, err := base62.Random(10)
	if err != nil {
		t.Fatalf("unable to generate id: %v", err)
	}
	compositeResource := ModelWithCompositeKey{
		KeyOne: id,
		KeyTwo: id2,
	}
	if err := DB.Create(&compositeResource).Error; err != nil {
		t.Errorf("errors happened when create: %v", err)
	}
	if res := DB.Delete(&compositeResource); res.Error != nil || res.RowsAffected != 1 {
		t.Errorf("errors happened when delete: %v, affected: %v", res.Error, res.RowsAffected)
	}
}
