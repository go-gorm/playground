package v1

import "gorm.io/gorm"

func Migrate(m gorm.Migrator) error {
	return m.AutoMigrate(&User{}, &Company{}, &Language{}, &Pet{}, &Toy{}, &Account{})
}
