package main

import (
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
	var parId, semiParId *uint
	err = DB.Raw("select parent_id, semi_parent_id from children limit 1").Row().Scan(&parId, &semiParId)
	if err != nil {
		t.Error("Doesn't expect error from raw query, but was: ", err)
	}
	if parId == nil {
		t.Error("Expect parent id not nil; but retrieved nil")
	}
	if semiParId == nil {
		t.Error("Expect semi parent id not nil; but retrieved nil")
	}
}
