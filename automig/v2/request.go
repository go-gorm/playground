package automigv2

import "gorm.io/gorm"

// Request is a test struct to validate automigration of indexes with Postgres.
// In this second version, the Request modelthe unique index for name has been removed
// So when migrating from Request.V1 to RequestV2 we should be able to add new entries
// with the same name
type Request struct {
	gorm.Model
	Name string `json:"name"`
}
