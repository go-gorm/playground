package main

import (
	"testing"
	"time"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: mysql

type Blueprint struct {
	Id                  uint64
	Name                string    `gorm:"size:255;not null;index"`
	CreatedAt           time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt           time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	BlueprintProperties []*BlueprintProperty
}

type Property struct {
	Id        uint64
	Name      string    `gorm:"size:255;not null;index"`
	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

type BlueprintProperty struct {
	Id          uint64
	BlueprintId uint64     `gorm:"not null;index:idx_blueprint_properties_blueprint_id_property_id"`
	PropertyId  uint64     `gorm:"not null;index:idx_blueprint_properties_blueprint_id_property_id"`
	Value       string     `gorm:"size:255;not null"`
	CreatedAt   time.Time  `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time  `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	Blueprint   *Blueprint `gorm:"foreignkey:BlueprintId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Property    *Property  `gorm:"foreignkey:PropertyId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type BlueprintPropertyView struct {
	BlueprintId uint64
	PropertyId  uint64
	Name        string
	Value       string
}

func (BlueprintPropertyView) TableName() string {
	return "blueprint_properties_view"
}

type Product struct {
	Id          uint64
	BlueprintId uint64     `gorm:"not null"`
	Quantity    uint16     `gorm:"size:16;not null"`
	CreatedAt   time.Time  `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time  `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	Blueprint   *Blueprint `gorm:"foreignkey:BlueprintId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type ProductView struct {
	Id                      uint64
	Quantity                int
	BlueprintId             uint64
	Name                    string
	BlueprintPropertiesView []*BlueprintPropertyView `gorm:"foreignKey:BlueprintId"`
}

func (ProductView) TableName() string {
	return "products_view"
}

func migrate() {
	DB.AutoMigrate(&Blueprint{})
	DB.AutoMigrate(&Property{})
	DB.AutoMigrate(&BlueprintProperty{})
	DB.AutoMigrate(&Product{})
	blueprintPropertyViewQuery := `
CREATE VIEW blueprint_properties_view AS
SELECT
  blueprint_properties.blueprint_id,
  blueprint_properties.property_id,
  properties.name,
  blueprint_properties.value
FROM blueprint_properties
INNER JOIN properties ON properties.id = blueprint_properties.property_id
  `
	DB.Exec(blueprintPropertyViewQuery)
	productViewQuery := `
CREATE VIEW products_view AS
SELECT
  products.id,
  products.quantity,
  products.blueprint_id,
  blueprints.name
FROM products
INNER JOIN blueprints ON blueprints.id = products.blueprint_id
  `
	DB.Exec(productViewQuery)
}

func dropAll() {
	DB.Exec("DROP VIEW products_view")
	DB.Exec("DROP VIEW blueprint_properties_view")
	DB.Exec("DROP TABLE blueprint_properties")
	DB.Exec("DROP TABLE products")
	DB.Exec("DROP TABLE properties")
	DB.Exec("DROP TABLE blueprints")
}

func TestGORM(t *testing.T) {
	dropAll()
	migrate()

	blueprint := Blueprint{Name: "blueprint_name"}
	DB.Create(&blueprint)

	property := Property{Name: "property_name"}
	DB.Create(&property)

	blueprintPropertySize := 10
	blueprintProperty := BlueprintProperty{BlueprintId: blueprint.Id, PropertyId: property.Id, Value: "prop_value"}
	for i := 0; i < blueprintPropertySize; i++ {
		blueprintProperty.Id = 0
		DB.Create(&blueprintProperty)
	}
	product := Product{BlueprintId: blueprint.Id, Quantity: 2}
	DB.Create(&product)
	product = Product{BlueprintId: blueprint.Id, Quantity: 3}
	DB.Create(&product)

	productsView := make([]ProductView, 0)
	DB.Preload("BlueprintPropertiesView").Find(&productsView)

	if len(productsView[0].BlueprintPropertiesView) != blueprintPropertySize {
		t.Errorf("Should be %d", blueprintPropertySize)
	}

	if len(productsView[1].BlueprintPropertiesView) != blueprintPropertySize {
		t.Errorf("Should be %d", blueprintPropertySize)
	}
}
