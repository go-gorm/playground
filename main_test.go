package main

import (
	"testing"

	"gorm.io/gorm/clause"

	"github.com/stretchr/testify/assert"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type M1 struct {
	ID string `gorm:"column:id; primaryKey; not null"`

	M2   *M2
	M2ID *string `gorm:"column:m2_id"`
}

type M2 struct {
	ID string `gorm:"column:id; primaryKey; not null"`

	M1   *M1
	M1ID *string `gorm:"column:m1_id"`
}

func TestGORM(t *testing.T) {
	DB.Migrator().DropTable(&M1{}, &M2{})
	DB.Migrator().AutoMigrate(&M1{}, &M2{})

	// Create

	m1 := M1{
		ID: "m1",
	}
	m2 := M2{
		ID: "m2",
	}

	assert.NoError(t, DB.Create(&m1).Error)
	assert.NoError(t, DB.Create(&m2).Error)

	// Set foreign keys
	m1.M2ID = &m2.ID
	m2.M1ID = &m1.ID
	assert.NoError(t, DB.Model(&m1).Updates(&m1).Error)
	assert.NoError(t, DB.Model(&m2).Updates(&m2).Error)
	m1.M2 = &m2
	m2.M1 = &m1

	var findM1 M1
	assert.NoError(t, DB.First(&findM1, M1{ID: m1.ID}).Error)
	assert.Equal(t, findM1.ID, m1.ID)
	assert.Equal(t, *findM1.M2ID, m2.ID)
	assert.Nil(t, findM1.M2) // No preload

	var findM2 M2
	assert.NoError(t, DB.First(&findM2, M2{ID: m2.ID}).Error)
	assert.Equal(t, findM2.ID, m2.ID)
	assert.Equal(t, *findM2.M1ID, m1.ID)
	assert.Nil(t, findM2.M1) // No preload

	// Now using preload all

	assert.NoError(t, DB.Preload(clause.Associations).First(&findM1, M1{ID: m1.ID}).Error)
	assert.Equal(t, findM1.ID, m1.ID)
	assert.Equal(t, *findM1.M2ID, m2.ID)
	assert.Equal(t, findM1.M2.ID, m2.ID) // With preload

	assert.NoError(t, DB.Preload(clause.Associations).First(&findM2, M2{ID: m2.ID}).Error)
	assert.Equal(t, findM2.ID, m2.ID)
	assert.Equal(t, *findM2.M1ID, m1.ID)
	assert.Equal(t, findM2.M1.ID, m1.ID) // With preload

	// Break things by adding new database entries
	{
		m1New := M1{
			ID: "m1-new",
		}
		m2New := &M2{
			ID: "m2-new",
		}

		assert.NoError(t, DB.Create(&m1New).Error)
		assert.NoError(t, DB.Create(&m2New).Error)

		// Set foreign keys
		m1New.M2ID = &m2.ID
		m2New.M1ID = &m1.ID
		assert.NoError(t, DB.Model(&m1New).Updates(&m1New).Error)
		assert.NoError(t, DB.Model(&m2New).Updates(&m2New).Error)
		m1New.M2 = &m2
		m2New.M1 = &m1
	}

	// Now the same query as above yields different results

	assert.NoError(t, DB.Preload(clause.Associations).First(&findM1, M1{ID: m1.ID}).Error)
	assert.Equal(t, findM1.ID, m1.ID)
	assert.Equal(t, *findM1.M2ID, m2.ID)
	assert.Equal(t, findM1.M2.ID, m2.ID) // With preload

	assert.NoError(t, DB.Preload(clause.Associations).First(&findM2, M2{ID: m2.ID}).Error)
	assert.Equal(t, findM2.ID, m2.ID)
	assert.Equal(t, *findM2.M1ID, m1.ID)
	assert.Equal(t, findM2.M1.ID, m1.ID) // With preload

	// Optional: Break even more things
	{
		m1New2 := M1{
			ID: "m1-new2",
		}
		m2New2 := &M2{
			ID: "m2-new2",
		}

		assert.NoError(t, DB.Create(&m1New2).Error)
		assert.NoError(t, DB.Create(&m2New2).Error)

		// Set foreign keys
		m1New2.M2ID = &m2.ID
		m2New2.M1ID = &m1.ID
		m1New2.M2 = &m2
		m2New2.M1 = &m1
		assert.NoError(t, DB.Model(&m1New2).Updates(&m1New2).Error)
		assert.NoError(t, DB.Model(&m2New2).Updates(&m2New2).Error)
	}

	// As we not only set the foreign ID but also the foreign object before
	// updating above, the find operation fails now (which is a bug).
	// The two issued queries don't make sense due to the duplicate "id" condition:
	// SELECT * FROM `m1` WHERE `m1`.`id` = "m1-new2" AND `m1`.`id` = "m1" ORDER BY `m1`.`id` LIMIT 1
	// SELECT * FROM `m2` WHERE `m2`.`id` = "m2-new2" AND `m2`.`id` = "m2" ORDER BY `m2`.`id` LIMIT 1

	assert.NoError(t, DB.First(&findM1, M1{ID: m1.ID}).Error)

	assert.NoError(t, DB.First(&findM2, M2{ID: m2.ID}).Error)
}
