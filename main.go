package main

import "gorm.io/gorm"

// Preload helper to create a scope for preloading associations.
func Preload(query string, args ...interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload(query, args...)
	}
}

// QueryColumnsFilter provides a scope to filter columns.
// This is the scope containing the OR condition that triggers the bug.
// It should select columns that are either system columns OR are not hidden.
func QueryColumnsFilter() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// The condition is: (is_system = true OR is_hidden = false)
		return db.Where(db.Where("is_system = ?", true).Or("is_hidden = ?", false))
	}
}