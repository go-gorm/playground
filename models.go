package main

type TableOne struct {
	ID uint `gorm:"primaryKey;autoIncrement"`

	TableTwoID uint
	TableTwo   TableTwo
}

func (to TableOne) TableName() string {
	return "my_schema.table_oneeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
}

type TableTwo struct {
	ID uint `gorm:"primaryKey;autoIncrement"`
}

func (tt TableTwo) TableName() string {
	return "my_schema.table_two"
}
