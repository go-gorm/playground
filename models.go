package main

// User has one `Account` (has one), many `Pets` (has many) and `Toys` (has many - polymorphic)
// He works in a Company (belongs to), he has a Manager (belongs to - single-table), and also managed a Team (has many - single-table)
// He speaks many languages (many to many) and has many friends (many to many - single-table)
// His pet also has one Toy (has one - polymorphic)
type User struct {
	ID       string
	Revision int `gorm:"primaryKey"`
	Name     string
	Embed    *Embed `gorm:"embedded;embeddedPrefix:embed_"`
}
type Embed struct {
	FieldName string
}
