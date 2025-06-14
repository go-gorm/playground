package main

import (
	"fmt"

	"gorm.io/gorm"
)

func main() {
	fmt.Println("vim-go")
}

func (m *Man) update(data interface{}) error {
	return DB.Set("data", data).Model(m).Where("id = ?", m.Id).Updates(data).Error
}

func (m *Man) BeforeUpdate(tx *gorm.DB) error {
	if !tx.Statement.Changed("age") {
		fmt.Println("no")
		return nil
	}
	fmt.Println("yes")
	return nil
}
