package main

type User struct {
	Id   uint64 `gorm:"column:id;type:bigserial;primary_key;unique_index"`
	Name string `gorm:"column:name;type:text"`
}