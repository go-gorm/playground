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
	if err := DB.First(&result, user.Id).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	product := Product{Name: "product", MyModel: MyModel{CreatedBy: user.Id}, OwnerId: user.Id}
	store := Store{Name: "store", Products: []Product{product}}
	DB.Create(&store)

	var result2 []Store
	if err := DB.Preload("Products.Owner").Find(&result2).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	if result2[0].Products[0].Owner.Id != user.Id && result2[0].Products[0].Owner.Id != 0 {
		t.Errorf("Failed, expected Owner ID to be %v, got %v", user.Id, result2[0].Products[0].Owner.Id)
	}
}
