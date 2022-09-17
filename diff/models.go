package diff

import (
	"time"

	"gorm.io/gorm"
)

type GameUser struct {
	gorm.Model
	Nick         string `gorm:"uniqueIndex;size:255"`
	Clan         string `gorm:"size:255"`
	ClanUrl      string `gorm:"size:255"`
	Banned       *bool
	RegisterDate time.Time
	Title        string `gorm:"size:255"`
	Level        int
	StatAb       UserStat `gorm:"embedded;embeddedPrefix:stat_ab_"`
	StatRb       UserStat `gorm:"embedded;embeddedPrefix:stat_rb_"`
	StatSb       UserStat `gorm:"embedded;embeddedPrefix:stat_sb_"`

	GroundRateAb GroundRate `gorm:"embedded;embeddedPrefix:rate_ground_ab_"`
	GroundRateRb GroundRate `gorm:"embedded;embeddedPrefix:rate_ground_rb_"`
	GroundRateSb GroundRate `gorm:"embedded;embeddedPrefix:rate_ground_sb_"`

	AviationRateAb AviationRate `gorm:"embedded;embeddedPrefix:rate_aviation_ab_"`
	AviationRateRb AviationRate `gorm:"embedded;embeddedPrefix:rate_aviation_rb_"`
	AviationRateSb AviationRate `gorm:"embedded;embeddedPrefix:rate_aviation_sb_"`

	FleetRateAb FleetRate `gorm:"embedded;embeddedPrefix:rate_fleet_ab_"`
	FleetRateRb FleetRate `gorm:"embedded;embeddedPrefix:rate_fleet_rb_"`
	FleetRateSb FleetRate `gorm:"embedded;embeddedPrefix:rate_fleet_sb_"`

	TsABRate float64
	TsRBRate float64
	TsSBRate float64
	AsABRate float64
	AsRBRate float64
	AsSBRate float64
}

type UserStat struct {
	TotalMission         int
	WinRate              float64
	GroundDestroyCount   int
	FleetDestroyCount    int
	GameTime             string
	AviationDestroyCount int
	WinCount             int
	SliverEagleEarned    int64
	DeadCount            int
}

type GroundRate struct {
	GameCount              int
	GroundVehicleGameCount int
	TDGameCount            int
	HTGameCount            int
	SPAAGameCount          int
	GameTime               string
	GroundVehicleGameTime  string
	TDGameTime             string
	HTGameTime             string
	SPAAGameTime           string
	TotalDestroyCount      int
	AviationDestroyCount   int
	GroundDestroyCount     int
	FleetDestroyCount      int
}

type AviationRate struct {
	GameCount            int
	FighterGameCount     int
	BomberGameCount      int
	AttackerGameCount    int
	GameTime             string
	FighterGameTime      string
	BomberGameTime       string
	AttackerGameTime     string
	TotalDestroyCount    int
	AviationDestroyCount int
	GroundDestroyCount   int
	FleetDestroyCount    int
}

type FleetRate struct {
	GameCount               int
	FleetGameCount          int
	TorpedoBoatGameCount    int
	GunboatGameCount        int
	TorpedoGunboatGameCount int
	SubmarineHuntGameCount  int
	DestroyerGameCount      int
	NavyBargeGameCount      int
	GameTime                string
	FleetGameTime           string
	TorpedoBoatGameTime     string
	GunboatGameTime         string
	TorpedoGunboatGameTime  string
	SubmarineHuntGameTime   string
	DestroyerGameTime       string
	NavyBargeGameTime       string
	TotalDestroyCount       int
	AviationDestroyCount    int
	GroundDestroyCount      int
	FleetDestroyCount       int
}
