package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite

func TestGORM(t *testing.T) {
	DB.AutoMigrate(&Coupon{}, &CouponProduct{})

	err := DB.
		Session(&gorm.Session{FullSaveAssociations: true}).
		Create(&Coupon{
			ID: "coupon1",
			AppliesToProduct: []*CouponProduct{
				{
					CouponId:  "coupon1",
					ProductId: "prod1",
				},
			},
			AmountOff:  10,
			PercentOff: 0.0,
		}).Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}

type Coupon struct {
	ID               string           `gorm:"primarykey; size:255"`
	AppliesToProduct []*CouponProduct `gorm:"foreignKey:CouponId;constraint:OnDelete:CASCADE"`
	AmountOff        uint32           `gorm:"amount_off"`
	PercentOff       float32          `gorm:"percent_off"`
}

type CouponProduct struct {
	CouponId  string `gorm:"primarykey; size:255"`
	ProductId string `gorm:"primarykey; size:255"`
}
