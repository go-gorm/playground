package main

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/clause"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestSave(t *testing.T) {
	ch := Child{
		EvalField: "ch",
	}
	semiParent := SemiParent{
		EvalField: "sp",
		Children:  []*Child{&ch},
	}
	parent := Parent{
		EvalField:   "p",
		Children:    []*Child{&ch},
		SemiParents: []*SemiParent{&semiParent},
	}
	err := DB.Debug().Create(&parent).Error
	//one of the queries INSERT INTO "children" ("eval_field","parent_id","semi_parent_id") VALUES ('ch',1,NULL)... - but both foreign keys must not be null
	assert.Nil(t, err)
	//Here both children foreign keys populated in the model
	assert.NotNil(t, ch.ParentID)
	assert.NotNil(t, ch.SemiParentID)
	//check passes, but only one foreign key populated in DB

	//So get saved entity from db and validate is all relations exists
	dbPar := new(Parent)
	err = DB.Debug().
		Preload(clause.Associations).
		Find(&dbPar, "parents.id = ?", parent.ID).Error
	assert.Nil(t, err)
	assert.NotNil(t, dbPar.SemiParents)
	assert.NotNil(t, dbPar.Children)
	assert.NotEmpty(t, dbPar.SemiParents, dbPar.Children)   //both root 'has many' collections not empty
	assert.True(t, len(dbPar.SemiParents) == 1)             //one semi parent found
	assert.True(t, len(dbPar.Children) == 1)                //one root child found
	assert.True(t, len(dbPar.SemiParents[0].Children) == 1) //Fail here! cause foreign key wasn't populated
}
