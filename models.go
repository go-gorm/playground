package main

type A struct {
	ID uint
}

type B struct {
	A `gorm:"embedded"`
}

type C struct {
	ID uint
	B  B
}
