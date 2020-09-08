package main

import (
	"github.com/oklog/ulid/v2"
)

type One struct {
	ID ulid.ULID `gorm:"primaryKey;type:varbinary(255)"`

	Name string

	Two   *Two
	TwoID ulid.ULID
}

type Two struct {
	ID ulid.ULID `gorm:"primaryKey;type:varbinary(255)"`
}
