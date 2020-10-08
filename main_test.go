package main

import (
	"os"
	"reflect"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	// to have the foreign keys on sqlite, because its needed for the test
	if os.Getenv("GORM_DIALECT") == "sqlite" {
		err := DB.Exec("PRAGMA foreign_keys = ON;").Error
		if err != nil {
			t.Fatal(err)
		}
	}

	// migrate the model for this test
	err := DB.AutoMigrate(&Root{}, &Branch{}, &Leaf{})
	if err != nil {
		t.Fatal(err)
	}

	// three different models which should all work:
	// single: one root with one branch 'b1' with one leaf 'l1'
	// two leaves: one root with one branch '1' with two leaves 'l1' and 'l2'
	// two branches one root with two branches 'b1' and 'b2' with one leaf 'l1' each
	var tests = map[string]*Root{
		"single": {
			Branches: []*Branch{
				{
					Name: "b1",
					Leaves: []*Leaf{
						{Name: "l1"},
					},
				},
			},
		},

		"two leaves": {
			Branches: []*Branch{
				{
					Name: "b1",
					Leaves: []*Leaf{
						{Name: "l1"},
						{Name: "l2"},
					},
				},
			},
		},

		"two branches": {
			Branches: []*Branch{
				{
					Name: "b1",
					Leaves: []*Leaf{
						{Name: "l1"},
					},
				},
				{
					Name: "b2",
					Leaves: []*Leaf{
						{Name: "l2"},
					},
				},
			},
		},
	}

	// parameterized test for each model
	for name, model := range tests {
		t.Run(name, func(t *testing.T) {
			// create the model for the test
			err := DB.Create(model).Error
			if err != nil {
				t.Fatal(err)
			}

			// now reload the modal with the preloading for the complete tree
			var loaded Root
			err = DB.Debug().Preload("Branches.Leaves").Take(&loaded, model.ID).Error
			if err != nil {
				t.Fatal(err)
			}

			// check if in and out are the same
			if !reflect.DeepEqual(model, &loaded) {
				t.Error("not equal")
			}
		})
	}
}

// model classes for this test
type (
	// Root is the root table with an autoincrement id
	// it references Branch by branch's composite primary key
	Root struct {
		ID       uint64    `gorm:"primaryKey;autoIncrement"`
		Branches []*Branch `gorm:"foreignKey:RootId;constraint:OnUpdate:RESTRICT,OnDelete:CASCADE;"`
	}

	// Branch is has a composite primary key derived from the ID of its Root and its own Name
	// it references Leaf by leaf's composite primary key
	Branch struct {
		RootId uint64  `gorm:"primaryKey"`
		Name   string  `gorm:"primaryKey;size:100"`
		Leaves []*Leaf `gorm:"foreignKey:RootId,BranchName;constraint:OnUpdate:RESTRICT,OnDelete:CASCADE;"`
	}

	// Leaf is has a composite primary key derived from the ID of its Root, Name of its Branch and its own Name
	Leaf struct {
		RootId     uint64 `gorm:"primaryKey"`
		BranchName string `gorm:"primaryKey;size:100"`
		Name       string `gorm:"primaryKey;size:100"`
	}
)
