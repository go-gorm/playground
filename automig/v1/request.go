package automigv1

import "gorm.io/gorm"

// Request is a test struct to validate automigration of indexes with Postgres.
// In this first version, the Request model has a unique Index for Names
type Request struct {
	gorm.Model
	Name string `json:"name" gorm:"unique"`
}
