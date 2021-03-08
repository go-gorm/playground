package main

import (
	"fmt"
	"testing"
	"time"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	name := "unique-user" + time.Now().Format(time.RFC3339)
	user := &User{Name: name}

	// Observed postgres SQL:
	// INSERT INTO "users" ("created_at","updated_at","deleted_at","name","age","birthday","company_id","manager_id","active") VALUES ('2021-03-08 12:32:59.863','2021-03-08 12:32:59.863',NULL,'',0,NULL,NULL,NULL,false) RETURNING "id"
	// A record is inserted, but the name is NULL
	var createResult User
	// A double pointer triggers this behavior
	if result := DB.FirstOrCreate(&createResult, &user); result.Error != nil {
		t.Errorf("Failed, got error: %v", result.Error)
	} else {
		fmt.Printf("Rows affected %d\n", result.RowsAffected)
	}

	fmt.Printf("Condition param user ID (expect 0/unchaged) %d\n", user.ID)
	fmt.Printf("Result user ID %d\n", createResult.ID)

	// If no error from Create and a positive ID, the record should have been created and we should be able to find it.
	if createResult.ID > 0 {
		var findResult User
		if err := DB.Model(&User{}).Where("name = ?", name).Find(&findResult).Error; err != nil {
			t.Errorf("Failed, got error: %v", err)
		}

		if findResult.ID != createResult.ID {
			t.Fatalf("User from create and user from find to not match. Create ID %d, find ID %d", findResult.ID, createResult.ID)
		}
	}
}

func TestRecordViolatesConstraint(t *testing.T) {
	type Constraint struct {
		gorm.Model
		Count int
	}
	DB.AutoMigrate(&Constraint{})
	DB.Exec(`ALTER TABLE constraints ADD CONSTRAINT count_check CHECK (count > 0);`)

	user := &Constraint{}
	var createResult Constraint
	// A double pointer triggers this behavior
	if result := DB.FirstOrCreate(&createResult, &user); result.Error != nil {
		// This should error but does not.
		// If plain `Create` is used there's an error. If a single pointer with FirstOrCreate is used a panic occurs.
		t.Errorf("Failed, got error: %v", result.Error)
	} else if result.RowsAffected != 1 {
		// No error so presumably created?
		t.Fatalf("Expected a row to have been created")
	} else {
		fmt.Printf("Rows %d\n", result.RowsAffected)
	}
}

func TestPlainCreate(t *testing.T) {
	name := "create-user" + time.Now().Format(time.RFC3339)
	user := &User{Name: name}

	// Even with a double pointer this generates the record.
	if err := DB.Create(&user).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	fmt.Printf("Result user ID %d\n", user.ID)

	// If no error from Create and a positive ID, the record should have been created and we should be able to find it.
	if user.ID > 0 {
		var findResult User
		if err := DB.Model(&User{}).Where("name = ?", name).Find(&findResult).Error; err != nil {
			t.Errorf("Failed, got error: %v", err)
		}

		if findResult.ID != user.ID {
			t.Fatalf("User from create and user from find to not match. Create ID %d, find ID %d", findResult.ID, user.ID)
		}
	}
}
