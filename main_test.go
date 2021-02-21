package main

import (
	"log"
	"os"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestAddUniqueConstraintOnExistingColumn(t *testing.T) {
	DB.Migrator().DropTable(&User1{})

	// Migrate User table without unique constraint
	if err := DB.AutoMigrate(&User1{}); err != nil {
		log.Printf("Failed to auto migrate, but got error %v\n", err)
		os.Exit(1)
	}

	exist := DB.Migrator().HasConstraint(&User2{}, "users_email_key")
	if exist {
		t.Errorf("constraint exist. it should not")
	}

	// Then, we add a unique constraint
	if err := DB.AutoMigrate(&User2{}); err != nil {
		log.Printf("Failed to auto migrate, but got error %v\n", err)
		os.Exit(1)
	}

	exist = DB.Migrator().HasConstraint(&User2{}, "users_email_key")
	if !exist {
		t.Errorf("constraint does not exist. it should")
	}
}

func TestUniqueConstraintOnNewTable(t *testing.T) {
	DB.Migrator().DropTable(&SomeNewTable{})

	if err := DB.AutoMigrate(&SomeNewTable{}); err != nil {
		log.Printf("Failed to auto migrate, but got error %v\n", err)
		os.Exit(1)
	}
	exist := DB.Migrator().HasConstraint(&SomeNewTable{}, "some_new_tables_email_key")
	if exist {
		t.Errorf("constraint does not exist. it should")
	}
}
