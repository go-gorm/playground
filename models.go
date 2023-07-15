package main

type MyModel struct {
	Id        uint `gorm:"primaryKey"`
	CreatedBy uint
}

type User struct {
	MyModel
	Name string
}

type Store struct {
	MyModel
	Name     string
	Products []Product `gorm:"foreignKey:StoreId"`
}

type Product struct {
	MyModel
	StoreId uint
	Name    string
	OwnerId uint
	// Owner   User `gorm:"foreignKey:OwnerId;references:Id"` // this works
	Owner User `gorm:"foreignKey:CreatedBy;references:Id"` // this not work, CreatedBy here belongs to Product, not User
}
