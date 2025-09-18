package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: mysql

func TestGORM(t *testing.T) {
	type GormBugTest struct {
		IDS []string `gorm:"serializer:json;column:ids;type:varchar(1000);not null;default:'名称';comment:ids" json:"ids"`
	}

	for i := 0; i < 3; i++ {
		// If a field in a database table uses the serializer tag and AutoMigrate is used,
		// then even if there are no changes to this field
		// ,AutoMigrate will still perform an ALTER on the table every time.

		// This will alter the ids field of the GormBugTest table four times,
		// even though the field hasn't changed at all.
		err := DB.AutoMigrate(GormBugTest{})
		if err != nil {
			t.Errorf("Failed, got error: %v", err)
		}
	}

	t.Log("This will alter the ids field of the GormBugTest table four times,even though the field hasn't changed at all.")
	t.Log("这里会alter四次GormBugTest表的ids字段 虽然字段没有任何变化")
}
