package main

// User has one `Account` (has one), many `Pets` (has many) and `Toys` (has many - polymorphic)
// He works in a Company (belongs to), he has a Manager (belongs to - single-table), and also managed a Team (has many - single-table)
// He speaks many languages (many to many) and has many friends (many to many - single-table)
// His pet also has one Toy (has one - polymorphic)
type User struct {
	ID        *int `gorm:"primarykey"`
	Name      string
	CompanyID *int
	Languages []Language `gorm:"many2many:UserSpeak;ForeignKey:ID;References:Code,CompanyID;joinForeignKey:user_id;JoinReferences:code,company_id;"`
}

type Language struct {
	ID        *int `gorm:"primarykey"`
	Code      int  `gorm:"index:idx_lan"`
	CompanyID *int `gorm:"index:idx_lan"`
	Name      string
}
