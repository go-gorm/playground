package main

type Coupon struct {
	ID               int64            `gorm:"primarykey; size:255"`
	AppliesToProduct []*CouponProduct `gorm:"foreignKey:CouponId;constraint:OnDelete:CASCADE"`
}

type CouponProduct struct {
	CouponId  int64 `gorm:"primaryKey"`
	ProductId int   `gorm:"primaryKey"`
}
