package main

import (
	"fmt"
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, postgres

func TestGORM(t *testing.T) {
	type BatchMapUser struct {
		ID   uint
		Name string
		Age  int
	}
	if err := DB.AutoMigrate(&BatchMapUser{}); err != nil {
		t.Fatal(err)
	}
	for _, pointer := range []bool{false, true} {
		for _, configured := range []bool{false, true} {
			for _, size := range []struct{ count, batch int }{{1, 500}, {5, 2}} {
				t.Run(fmt.Sprintf("pointer=%t/configured=%t/count=%d", pointer, configured, size.count), func(t *testing.T) {
					if err := DB.Where("1 = 1").Delete(&BatchMapUser{}).Error; err != nil {
						t.Fatal(err)
					}
					rows := make([]map[string]interface{}, size.count)
					for i := range rows {
						rows[i] = map[string]interface{}{"Name": fmt.Sprintf("user%d", i), "Age": 18 + i}
					}
					var value interface{} = rows
					if pointer {
						value = &rows
					}
					tx := DB.Model(&BatchMapUser{})
					var result *gorm.DB
					if configured {
						result = tx.Session(&gorm.Session{CreateBatchSize: size.batch}).Create(value)
					} else {
						result = tx.CreateInBatches(value, size.batch)
					}
					if result.Error != nil {
						t.Fatal(result.Error)
					}
					if result.RowsAffected != int64(size.count) {
						t.Fatalf("rows affected=%d, want %d", result.RowsAffected, size.count)
					}
					var stored []BatchMapUser
					if err := DB.Order("id").Find(&stored).Error; err != nil {
						t.Fatal(err)
					}
					if len(stored) != size.count || len(rows) != size.count {
						t.Fatalf("stored=%d input=%d, want %d", len(stored), len(rows), size.count)
					}
					for i, user := range stored {
						if user.Name != fmt.Sprintf("user%d", i) || user.Age != 18+i || user.ID == 0 {
							t.Fatalf("row %d: %+v", i, user)
						}
					}
				})
			}
		}
	}
}

func TestBatchControls(t *testing.T) {
	type BatchControlUser struct {
		ID   uint
		Name string
		Age  int
	}
	if err := DB.AutoMigrate(&BatchControlUser{}); err != nil {
		t.Fatal(err)
	}
	for _, maps := range []bool{false, true} {
		for _, configured := range []bool{false, true} {
			t.Run(fmt.Sprintf("table_maps=%t/configured=%t", maps, configured), func(t *testing.T) {
				if err := DB.Where("1 = 1").Delete(&BatchControlUser{}).Error; err != nil {
					t.Fatal(err)
				}
				users := []BatchControlUser{{Name: "alice", Age: 18}, {Name: "bob", Age: 19}, {Name: "carol", Age: 20}}
				var value interface{} = &users
				tx := DB.Session(&gorm.Session{})
				if maps {
					value = []map[string]interface{}{{"name": "alice", "age": 18}, {"name": "bob", "age": 19}, {"name": "carol", "age": 20}}
					tx = tx.Table("batch_control_users") // No model, so no automatic RETURNING.
				}
				var result *gorm.DB
				if configured {
					result = tx.Session(&gorm.Session{CreateBatchSize: 2}).Create(value)
				} else {
					result = tx.CreateInBatches(value, 2)
				}
				if result.Error != nil || result.RowsAffected != 3 {
					t.Fatalf("error=%v rows=%d", result.Error, result.RowsAffected)
				}
				var stored []BatchControlUser
				if err := DB.Order("id").Find(&stored).Error; err != nil {
					t.Fatal(err)
				}
				if len(stored) != 3 {
					t.Fatalf("stored=%v", stored)
				}
				for i, u := range stored {
					if u.Name != users[i].Name || u.Age != users[i].Age || u.ID == 0 {
						t.Fatalf("row %d: %+v", i, u)
					}
					if !maps && u.ID != users[i].ID {
						t.Fatalf("row %d: ID not returned", i)
					}
				}
			})
		}
	}
}
