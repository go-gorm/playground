package main

type Country struct {
	Name string `gorm:"primaryKey"`
}

type Address struct {
	CountryName string
	Country     Country
}

type Org struct {
	ID       int
	Address1 Address `gorm:"embedded;embeddedPrefix:address1_"`
	Address2 Address `gorm:"embedded;embeddedPrefix:address2_"`
}
