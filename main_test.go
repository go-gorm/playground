package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	tests := []struct {
		name string
		id   int
	}{
		{name: "ok", id: 1},
		{name: "broken", id: 0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			coupon := Coupon{
				AppliesToProduct: []*CouponProduct{{CouponId: 1, ProductId: test.id}},
			}
			DB.Session(&gorm.Session{FullSaveAssociations: true}).Create(&coupon)

			var result Coupon
			if err := DB.Preload("AppliesToProduct").First(&result, "id = ?", coupon.ID).Error; err != nil {
				t.Errorf("Failed, got error: %v", err)
			}

			if len(result.AppliesToProduct) != 1 {
				t.FailNow()
			}
		})
	}
}
