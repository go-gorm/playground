package main

import (
	"fmt"
	"log"
	"time"
	"os"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"gorm-bug/types"
)

// SlingUser has a composite primary key (ID, SrcID)
type SlingUser struct {
	ID    int64      `gorm:"primaryKey;autoIncrement:false"`
	SrcID string     `gorm:"primaryKey"`
	User  types.User `gorm:"embedded;embeddedPrefix:user_"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (SlingUser) TableName() string {
	return "sling_users"
}

// SlingLocation has a composite primary key (ID, SrcID)
type SlingLocation struct {
	ID       int64          `gorm:"primaryKey;autoIncrement:false"`
	SrcID    string         `gorm:"primaryKey"`
	Location types.Location `gorm:"embedded;embeddedPrefix:location_"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (SlingLocation) TableName() string {
	return "sling_locations"
}

// SlingShift references SlingUser and SlingLocation via composite foreign keys
type SlingShift struct {
	ID         string      `gorm:"primaryKey"`
	SrcID      string      `gorm:"primaryKey"`
	UserID     int64       `gorm:"not null;index"`
	LocationID int64       `gorm:"index"`
	Shift      types.Shift `gorm:"embedded;embeddedPrefix:shift_"`

	User     *SlingUser     `gorm:"foreignKey:SrcID,UserID;references:SrcID,ID"`
	Location *SlingLocation `gorm:"foreignKey:SrcID,LocationID;references:SrcID,ID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (SlingShift) TableName() string {
	return "sling_shifts"
}

func main() {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	err = db.AutoMigrate(&SlingUser{}, &SlingLocation{}, &SlingShift{})
	if err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}

	srcID := "test-src-123"

	// Create test data
	user := &SlingUser{
		ID:    1001,
		SrcID: srcID,
		User: types.User{
			ID:        1001,
			Email:     "alice@example.com",
			FirstName: "Alice",
			LastName:  "Smith",
		},
	}
	location := &SlingLocation{
		ID:    2001,
		SrcID: srcID,
		Location: types.Location{
			ID:      2001,
			Name:    "Office A",
			Address: "123 Main St",
		},
	}
	shift := &SlingShift{
		ID:         "shift-001",
		SrcID:      srcID,
		UserID:     user.ID,
		LocationID: location.ID,
		Shift: types.Shift{
			ID:        "shift-001",
			Summary:   "Morning shift",
			Status:    "confirmed",
			StartTime: types.Time{Time: time.Now()},
			EndTime:   types.Time{Time: time.Now().Add(8 * time.Hour)},
			RawUser: types.User{
				ID:        1001,
				Email:     "alice@example.com",
				FirstName: "Alice",
				LastName:  "Smith",
			},
			Location: types.Ref{ID: 2001},
		},
	}

	db.Create(user)
	db.Create(location)
	db.Create(shift)

	fmt.Println("\n=== Loading shift with Preload ===")

	// Now try to load the shift with Preload
	var loadedShift SlingShift
	err = db.Preload("User").Preload("Location").
		Where("id = ? AND src_id = ?", "shift-001", srcID).
		First(&loadedShift).Error
	if err != nil {
		log.Fatalf("failed to load shift: %v", err)
	}

	fmt.Printf("\nShift ID: %s\n", loadedShift.ID)
	fmt.Printf("Shift UserID: %d\n", loadedShift.UserID)
	fmt.Printf("Shift LocationID: %d\n", loadedShift.LocationID)

	if loadedShift.User == nil {
		fmt.Println("ERROR: User is nil - Preload failed!")
		os.Exit(1)
	}

	fmt.Printf("User: %s %s (ID: %d)\n", loadedShift.User.User.FirstName, loadedShift.User.User.LastName, loadedShift.User.ID)

	if loadedShift.Location == nil {
		fmt.Println("ERROR: Location is nil - Preload failed!")
		os.Exit(1)
	}
	fmt.Printf("Location: %s (ID: %d)\n", loadedShift.Location.Location.Name, loadedShift.Location.ID)
}
