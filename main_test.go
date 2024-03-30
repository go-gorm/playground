package main

import (
	"fmt"
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

// I think second TestCase should pass but fails
func Test_primaryKey(t *testing.T) {
	tests := []struct {
		name     string
		prepare  func() error
		wantType string
		wantKey  string
	}{
		{
			name: "ok: TenantID to be Primary Key in init AutoMigrate",
			prepare: func() error {
				return DB.Table("tenant").AutoMigrate(&TenantWithPrimaryKey{})
			},
			wantType: "varchar(36)",
			wantKey:  "PRI",
		},
		{
			name: "should be ok but fails: TenantID to be Primary Key in second AutoMigrate",
			prepare: func() error {
				if err := DB.Table("tenant").AutoMigrate(&TenantWithoutPrimaryKey{}); err != nil {
					return err
				}
				return DB.Table("tenant").AutoMigrate(&TenantWithPrimaryKey{})
			},
			wantType: "varchar(36)",
			wantKey:  "PRI",
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

			type column struct {
				Type string
				Key  string
			}

			var c column
			if err := DB.Raw("SHOW COLUMNS FROM tenant").First(&c).Error; err != nil {
				t.Fatal(err)
			}

			fmt.Println(c)

			gotType, gotKey := c.Type, c.Key
			if gotType != tt.wantType {
				t.Fatalf("gotType = %s, wantType = %s", gotType, tt.wantType)
			}
			if gotKey != tt.wantKey {
				t.Fatalf("gotKey = %s, wantKey = %s", gotKey, tt.wantKey)
			}
		})
	}
}
