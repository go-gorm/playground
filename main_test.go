package main

import (
	"fmt"
	"gorm.io/gorm"
	"testing"

	uuid "github.com/satori/go.uuid"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type Application struct {
	ID        string `gorm:"primarykey"`
	Name      string
	Type      string
	Resources []Resource `json:"resources,omitempty" gorm:"many2many:application_resources;" faker:"-"`
}

type Resource struct {
	ID       string `gorm:"primarykey"`
	Name     string
	Type     string
	Packages []Package `json:"packages,omitempty" gorm:"many2many:resource_packages;" faker:"-"`
}

type Package struct {
	ID   string `gorm:"primarykey"`
	Name string
	Type string
}

func TestGORM(t *testing.T) {
	if err := DB.AutoMigrate(&Application{}, Resource{}, Package{}); err != nil {
		t.Fatalf("Failed to run auto migration: %v", err)
	}

	// Creates an application
	app := createApp(uuid.NewV4().String())
	if err := DB.Create(&app).Error; err != nil {
		t.Fatalf("failed to create application: %v", err)
	}

	if err := update(&app); err != nil {
		t.Fatalf("failed to update application: %v", err)
	}
}

func update(app *Application) error {
	// Updates application with 2 resources pointing to the same package and save to DB
	newApp := *app
	newApp.Resources = []Resource{createResource("1"), createResource("2")}
	if err := DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Save(newApp).Error; err != nil {

		// We failed here due to
		// ERROR: ON CONFLICT DO UPDATE command cannot affect row a second time (SQLSTATE 21000)
		//[2.167ms] [rows:0] INSERT INTO "packages" ("id","name","type") VALUES ('1','name','type'),('1','name','type') ON CONFLICT ("id") DO UPDATE SET "name"="excluded"."name","type"="excluded"."type"
		return fmt.Errorf("failed to save application: %v", err)
	}
	return nil
}

func createApp(id string) Application {
	return Application{
		ID:   id,
		Name: "name",
		Type: "type",
	}
}

func createResource(id string) Resource {
	return Resource{
		ID:       id,
		Name:     "name",
		Type:     "type",
		Packages: []Package{createPackage("1")},
	}
}

func createPackage(id string) Package {
	return Package{
		ID:   id,
		Name: "name",
		Type: "type",
	}
}
