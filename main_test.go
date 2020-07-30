package main

import (
	"testing"
	"github.com/gofrs/uuid"
)


type Service struct {
	ID uuid.UUID `json:"id,omitempty" gorm:"type:uuid;primary_key;"`

	// ID of a service group the service belongs to
	ServiceGroupID uuid.UUID `json:"service_group_id" gorm:"type:uuid;not null"`

	// ID of a host the service belongs to
	HostID uuid.UUID `json:"host_id" gorm:"type:uuid;not null"`
}

type Host struct {
	ID uuid.UUID `json:"id,omitempty" gorm:"type:uuid;primary_key;"`

	Services []*Service `json:"services,omitempty" gorm:"foreignkey:HostID; constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
}

type ServiceGroup struct {
	ID uuid.UUID `json:"id,omitempty" gorm:"type:uuid;primary_key;"`

	Services []*Service `json:"services,omitempty" gorm:"foreignkey:ServiceGroupID; constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
}

func TestGORM(t *testing.T) {
	err = DB.AutoMigrate(&Host{})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&Service{})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&ServiceGroup{})
	if err != nil {
		panic(err)
	}
	
	DB.Create(&Host{ID: 			uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111")}) //Create Host
	DB.Create(&ServiceGroup{ID: 		uuid.FromStringOrNil("55555555-5555-5555-5555-555555555555")}) //Create ServiceGroup
	DB.Create(&Service{ID: 			uuid.FromStringOrNil("33333333-3333-3333-3333-333333333333"),
			   HostID: 		uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"),
			   ServiceGroupID:	uuid.FromStringOrNil("55555555-5555-5555-5555-555555555555")
			  }) //Create a Valid service that belongs to both Host and ServiceGroup
	DB.Create(&Service{ID: 			uuid.FromStringOrNil("66666666-6666-6666-6666-666666666666"),
			   HostID: 		uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111")
			  }) //Create a another service that belongs to both Host BUT NOT ServiceGroup, this should result into an error
	
	var count int64
	DB.Table("rounds").Count(&count)
	if count != 1 {
		panic("Unexpected number of items in the table!")
	}
	
}
