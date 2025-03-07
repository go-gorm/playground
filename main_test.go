package main

import (
	"testing"
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
}

func TestDeletion(t *testing.T) {
	if DB.Migrator().HasTable(&Subject{}) {
		if err := DB.Migrator().DropTable(&Subject{}); err != nil {
			t.Logf("Failed to drop table, got error %v\n", err)
			t.Fail()
		}
	}
	if err := DB.AutoMigrate(&Subject{}); err != nil {
		t.Logf("Failed to auto migrate, but got error %v\n", err)
		t.Fail()
	}
	// Load data in
	if err := CreateSubjectFixtures(); err != nil {
		t.Logf("Failed to load data, but got error %v\n", err)
		t.Fail()
	}

	t.Run("structdelete", func(t *testing.T) {
		structDelete := Subject{}
		if err := DB.Debug().Model(&Subject{}).Where("subj_src = ?", "structdelete").First(&structDelete).Error; err != nil {
			t.Logf("Failed to first struct record, but got error %v\n", err)
			t.FailNow()
		}
		if err := DB.Debug().Delete(&structDelete).Error; err != nil {
			t.Logf("Failed to delete struct record, but got error %v\n", err)
			t.Fail()
		}
	})

	t.Run("structdeletebyid", func(t *testing.T) {
		structDelete := Subject{}
		if err := DB.Debug().Model(&Subject{}).Where("subj_src = ?", "structdeletebyid").First(&structDelete).Error; err != nil {
			t.Logf("Failed to first struct record, but got error %v\n", err)
			t.FailNow()
		}
		if err := DB.Debug().Delete(&Subject{}, structDelete.ID).Error; err != nil {
			t.Logf("Failed to delete structdeletebyid record, but got error %v\n", err)
			t.Fail()
		}
	})

	t.Run("singledeletebyid", func(t *testing.T) {
		structDelete := Subject{}
		if err := DB.Debug().Model(&Subject{}).Where("subj_src = ?", "singledeletebyid").First(&structDelete).Error; err != nil {
			t.Logf("Failed to first struct record, but got error %v\n", err)
			t.FailNow()
		}
		if err := DB.Debug().Where("id = ?", structDelete.ID).Delete(&Subject{}).Error; err != nil {
			t.Logf("Failed to delete singledeletebyid record, but got error %v\n", err)
			t.Fail()
		}
	})

	t.Run("singledelete", func(t *testing.T) {
		if err := DB.Debug().Where("subj_src = ?", "singledelete").Delete(&Subject{SubjSrc: "singledelete"}).Error; err != nil {
			t.Logf("Failed to delete single record, but got error %v\n", err)
			t.Fail()
		}
	})

	t.Run("bulkdelete", func(t *testing.T) {
		if err := DB.Debug().Where("subj_src = ?", "bulkdelete").Delete(&Subject{}).Error; err != nil {
			t.Logf("Failed to delete bulk record, but got error %v\n", err)
			t.Fail()
		}
	})
}
