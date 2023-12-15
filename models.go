package main

type Country1 struct {
	Name string `gorm:"primaryKey"`
}

type Country2 struct {
	CName string `gorm:"primaryKey"`
}

type Address1 struct {
	CountryName string
	Country     Country1
}

type Address2 struct {
	CName   string
	Country Country2 `gorm:"foreignKey:CName;references:CName"`
}

type Org struct {
	ID        int
	Adress1_1 Address1 `gorm:"embedded;embeddedPrefix:address1_1_"`
	Adress1_2 Address1 `gorm:"embedded;embeddedPrefix:address1_2_"`
	Adress2_1 Address2 `gorm:"embedded;embeddedPrefix:address2_1_"`
	Adress2_2 Address2 `gorm:"embedded;embeddedPrefix:address2_2_"`
}
