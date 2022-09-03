package main

import (
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
	if err != nil {
		t.Error("Expect nil error")
	}
	//Here both children foreign keys populated in the model
	if ch.ParentID == nil {
		t.Error("Expect parent foreign id not nil")
	}
	if ch.SemiParentID == nil {
		t.Error("Expect semi parent foreign id not nil")
	}
	//check passes, but only one foreign key populated in DB

	//So get saved entity from db and validate is all relations exists
	dbPar := new(Parent)
	err = DB.Debug().
		Preload(clause.Associations).
		Find(&dbPar, "parents.id = ?", parent.ID).Error
	if err != nil {
		t.Error("Expect nil error")
	}
	if dbPar.SemiParents == nil || len(dbPar.SemiParents) == 1 { //one semi parent found
		t.Error("Expect not nil and one element SemiParents collection")
	}
	if dbPar.Children == nil || len(dbPar.Children) == 1 { //one root child found
		t.Error("Expect not nil and one element  Children collection")
	}
	if len(dbPar.SemiParents[0].Children) != 1 { //Fail here! cause foreign key wasn't populated
		t.Error("Expect children collection populated with the one element; But found ", len(dbPar.SemiParents[0].Children))
	}
}
