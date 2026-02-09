package models

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type Layout struct {
	ID         uuid.UUID     `json:"id"`
	ProjectID  uuid.UUID     `json:"project_id" gorm:"<-:create"`
	PrivateFor uuid.NullUUID `json:"private_for"`
	Name       string        `json:"name"`
	CreatedAt  time.Time     `json:"created_at" gorm:"<-:create"`
	UpdatedAt  time.Time     `json:"updated_at"`
}
