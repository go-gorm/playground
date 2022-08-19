package main

import (
	"testing"
)

func TestGORM(t *testing.T) {
	var err error

	p1 := &Product{Code: "D1", Price: 100, Author: Author{
		Name: "chen",
	}}
	if err = DB.Create(p1).Error; err != nil {
		t.Error(err)
	}

	p2 := &Product{Code: "D2", Price: 100}
	if err = DB.Create(p2).Error; err != nil {
		t.Error(err)
	}
}
