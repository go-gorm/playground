package query

import "gorm.io/gen"

type CustomQuery interface {
	// SELECT * FROM @@table OFFSET 0 LIMIT 1
	List() (*gen.T, error)
}
