package main

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Subject struct {
	ID         uint   `gorm:"primarykey"`
	SubjSrc    string `gorm:"type:bytes;size:8;default:'';"`
	FileCount  int    `gorm:"default:0;"`
	PhotoCount int    `gorm:"default:0;"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (m *Subject) AfterDelete(tx *gorm.DB) (err error) {
	if err := tx.Model(m).Updates(map[string]interface{}{
		"FileCount":  0,
		"PhotoCount": 0,
	}).Error; err != nil {
		return fmt.Errorf("AfterDelete failed for m=%+v with err=%+v", m, err)
	}
	return nil
}

type SubjectMap map[string]Subject

func (m SubjectMap) Get(name string) Subject {
	if result, ok := m[name]; ok {
		return result
	}

	return Subject{}
}

func (m SubjectMap) Pointer(name string) *Subject {
	if result, ok := m[name]; ok {
		return &result
	}

	return &Subject{}
}

var SubjectFixtures = SubjectMap{
	"record1": Subject{SubjSrc: "bulkdelete",
		FileCount:  10,
		PhotoCount: 10,
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
		DeletedAt:  gorm.DeletedAt{},
	},
	"record2": Subject{SubjSrc: "bulkdelete",
		FileCount:  20,
		PhotoCount: 20,
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
		DeletedAt:  gorm.DeletedAt{},
	},
	"record3": Subject{SubjSrc: "bulkdelete",
		FileCount:  30,
		PhotoCount: 30,
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
		DeletedAt:  gorm.DeletedAt{},
	},
	"record4": Subject{SubjSrc: "bulkdelete",
		FileCount:  40,
		PhotoCount: 40,
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
		DeletedAt:  gorm.DeletedAt{},
	},
	"record5": Subject{SubjSrc: "bulkdelete",
		FileCount:  50,
		PhotoCount: 50,
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
		DeletedAt:  gorm.DeletedAt{},
	},
	"record-deleter": Subject{SubjSrc: "singledelete",
		FileCount:  60,
		PhotoCount: 60,
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
		DeletedAt:  gorm.DeletedAt{},
	},
	"record-structdeleter": Subject{SubjSrc: "structdelete",
		FileCount:  70,
		PhotoCount: 70,
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
		DeletedAt:  gorm.DeletedAt{},
	},
	"record-structdeletebyid": Subject{SubjSrc: "structdeletebyid",
		FileCount:  75,
		PhotoCount: 75,
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
		DeletedAt:  gorm.DeletedAt{},
	}, "record-singledeletebykey": Subject{SubjSrc: "singledeletebyid",
		FileCount:  80,
		PhotoCount: 80,
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
		DeletedAt:  gorm.DeletedAt{},
	},
	"record-keeper": Subject{SubjSrc: "keeper",
		FileCount:  999,
		PhotoCount: 999,
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
		DeletedAt:  gorm.DeletedAt{},
	}}

// CreateSubjectFixtures inserts known entities into the database for testing.
func CreateSubjectFixtures() (err error) {
	for _, entity := range SubjectFixtures {
		if err := DB.Create(&entity).Error; err != nil {
			return err
		}
	}
	return nil
}
