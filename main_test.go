package main

import (
	"fmt"
	"os"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}

// second TestCase should pass but fails
func Test_primaryKey(t *testing.T) {
	tests := []struct {
		name    string
		prepare func() error
	}{
		{
			name: "ok: TenantID to be Primary Key in init AutoMigrate",
			prepare: func() error {
				return DB.Table("tenant").AutoMigrate(&TenantWithPrimaryKey{})
			},
		},
		{
			name: "should be ok but fails: TenantID to be Primary Key in second AutoMigrate",
			prepare: func() error {
				if err := DB.Table("tenant").AutoMigrate(&TenantWithoutPrimaryKey{}); err != nil {
					return err
				}
				return DB.Table("tenant").AutoMigrate(&TenantWithPrimaryKey{})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.prepare(); err != nil {
				t.Fatal(err)
			}
			defer func() {
				if err := DB.Exec("DROP TABLE tenant").Error; err != nil {
					t.Fatal(err)
				}
			}()

			driver, found := os.LookupEnv("GORM_DIALECT")
			if !found {
				t.Fatalf("GORM_DIALECT should not be empty")
			}

			isPrimaryKey, err := isTenantIDPrimaryKey(driver)
			if err != nil {
				t.Fatal(err)
			}

			if !isPrimaryKey {
				t.Fatal("tenant_id should be the primary key")
			}
		})
	}
}

func isTenantIDPrimaryKey(driverName string) (bool, error) {
	switch driverName {
	case "mysql":
		type column struct {
			Key string
		}
		var c column
		if err := DB.Raw("SHOW COLUMNS FROM tenant").First(&c).Error; err != nil {
			return false, err
		}
		return c.Key == "PRI", nil
	case "postgres":
		type column struct {
			ConstraintName string
		}
		var c column
		if err := DB.Raw("SELECT constraint_name FROM information_schema.key_column_usage WHERE table_name = 'tenant'").First(&c).Error; err != nil {
			if err.Error() == "record not found" {
				return false, nil
			}
			return false, err
		}
		return c.ConstraintName == "tenant_pkey", nil
	case "sqlite":
		type column struct {
			PK bool
		}
		var c column
		if err := DB.Raw("PRAGMA table_info('tenant')").First(&c).Error; err != nil {
			return false, err
		}
		return c.PK, nil
	case "sqlserver":
		var hasPK bool
		if err := DB.Raw("SELECT OBJECTPROPERTY(object_id, 'TableHasPrimaryKey') FROM sys.tables WHERE name=tenant").First(&hasPK).Error; err != nil {
			return false, err
		}
		return hasPK, nil
	}
	return false, fmt.Errorf("invalid driver name %s", driverName)
}
