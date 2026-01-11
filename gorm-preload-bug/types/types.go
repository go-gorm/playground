package types

import "time"

// Time wraps time.Time (like sling.Time)
type Time struct {
	Time time.Time
}

// User represents a user (like sling.User embedded in Shift)
type User struct {
	ID        int64  `json:"id"`
	Deleted   bool   `json:"deleted"`
	Email     string `json:"email"`
	FirstName string `json:"name"`
	LastName  string `json:"lastname"`
}

// Ref is a reference to another entity (like sling.Ref)
type Ref struct {
	ID int64 `json:"id"`
}

// Location represents a location (like sling.Location)
type Location struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	State    string `json:"state"`
	Timezone string `json:"timezone"`
	Address  string `json:"address"`
}

// Shift represents shift details (like sling.Shift)
// This is embedded in the database model with a prefix
type Shift struct {
	ID            string `json:"id"`
	Summary       string `json:"summary"`
	Status        string `json:"status"`
	Type          string `json:"type"`
	FullDay       bool   `json:"fullDay"`
	OpenEnd       bool   `json:"openEnd"`
	StartTime     Time   `json:"dtstart" gorm:"embedded;embeddedPrefix:start_time_"`
	EndTime       Time   `json:"dtend" gorm:"embedded;embeddedPrefix:end_time_"`
	AssigneeNotes string `json:"assigneeNotes"`
	RawUser       User   `json:"user" gorm:"embedded;embeddedPrefix:raw_user_"`
	Location      Ref    `json:"location" gorm:"embedded;embeddedPrefix:location_"`
	Position      Ref    `json:"position" gorm:"embedded;embeddedPrefix:position_"`
	BreakDuration int64  `json:"breakDuration"`
	Available     bool   `json:"available"`
	Slots         int64  `json:"slots"`
}
