package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

// doc: https://gorm.io/docs/many_to_many.html#Override-Foreign-Key
type Companny struct {
	ID   int    `gorm:"primarykey"`
	Name string `gorm:"index:,unique"`
	Stus []Stu  `gorm:"many2many:stu_compannys;foreignKey:Name;joinForeignKey:CompannyName;References:Stuname;joinReferences:Stuname;"`
}
type Stu struct {
	ID        int       `gorm:"primarykey"`
	Stuname   string    `gorm:"index:,unique"`
	Compannies []Companny `gorm:"many2many:stu_compannys;"`
}

func TestMany2ManyUniqueKey(t *testing.T) {
    db := DB
    db.Debug().Migrator().DropTable(&Companny{}, &Stu{}, "stu_compannys")
	db.Debug().AutoMigrate(&Stu{})

	err := db.Debug().Create(&Stu{
		Stuname: "Alex3",
		ID:      3,
		Compannies: []Companny{
			{Name: "PKU1"}, {Name: "TSU2"},
		},
	})
	if err!=nil{
		t.Fatal(err)
	}

}
