package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {

	type Thing1 struct {
		gorm.Model
		Name string `gorm:"size:20"`
		One int
	}

	type Thing2 struct {
		gorm.Model
		Name string `gorm:"size:20"`
		Two int
	}

	type Thing3 struct {
		gorm.Model
		Name string `gorm:"size:20"`
		Three int
	}

	type Composite struct {
		*Thing1
		*Thing2
		*Thing3
	}

	if err := DB.AutoMigrate(&Thing1{}, &Thing2{}, &Thing3{}); err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	DB.Create(&Thing1{
		Name:  "Thing 1",
		One:   1,
	})
	DB.Create(&Thing2{
		Name:  "Thing 2",
		Two:   2,
	})
	DB.Create(&Thing3{
		Name:  "Thing 3",
		Three:   3,
	})

	t1 := new(Thing1)
	t2 := new(Thing2)
	t3 := new(Thing3)

	DB.First(t1)
	DB.First(t2)
	DB.First(t3)

	c := new(Composite)

	if err := DB.Raw("SELECT thing1.*, thing2.*, thing3.* FROM thing1 " +
		"INNER JOIN thing2 ON thing1.id = thing2.id " +
		"INNER JOIN thing3 ON thing1.id = thing3.id " +
		"WHERE thing1.id = 1").Scan(c).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if t1.ID != c.Thing1.ID {
		t.Errorf("Thing 1 ID. Wanted: %v - Got: %v", t1.ID, c.Thing1.ID)
	}
	if t1.Name != c.Thing1.Name {
		t.Errorf("Thing 1 Name. Wanted: %v - Got: %v", t1.Name, c.Thing1.Name)
	}
	if t1.CreatedAt != c.Thing1.CreatedAt {
		t.Errorf("Thing 1 Created. Wanted: %v - Got: %v", t1.CreatedAt, c.Thing1.CreatedAt)
	}
	if t1.One != c.Thing1.One {
		t.Errorf("Thing 1 One. Wanted: %v - Got: %v", t1.One, c.Thing1.One)
	}
	if t2.ID != c.Thing2.ID {
		t.Errorf("Thing 2 ID. Wanted: %v - Got: %v", t2.ID, c.Thing2.ID)
	}
	if t2.Name != c.Thing2.Name {
		t.Errorf("Thing 2 Name. Wanted: %v - Got: %v", t2.Name, c.Thing2.Name)
	}
	if t2.CreatedAt != c.Thing2.CreatedAt {
		t.Errorf("Thing 2 Created. Wanted: %v - Got: %v", t2.CreatedAt, c.Thing2.CreatedAt)
	}
	if t2.Two != c.Thing2.Two {
		t.Errorf("Thing 2 Two. Wanted: %v - Got: %v", t2.Two, c.Thing2.Two)
	}
	if t3.ID != c.Thing3.ID {
		t.Errorf("Thing 3 ID. Wanted: %v - Got: %v", t3.ID, c.Thing3.ID)
	}
	if t3.Name != c.Thing3.Name {
		t.Errorf("Thing 3 Name. Wanted: %v - Got: %v", t3.Name, c.Thing3.Name)
	}
	if t3.CreatedAt != c.Thing3.CreatedAt {
		t.Errorf("Thing 3 Created. Wanted: %v - Got: %v", t3.CreatedAt, c.Thing3.CreatedAt)
	}
	if t3.Three != c.Thing3.Three {
		t.Errorf("Thing 3 Three. Wanted: %v - Got: %v", t3.CreatedAt, c.Thing3.CreatedAt)
	}
}
