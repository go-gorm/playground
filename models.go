package main

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

// User has one `Account` (has one), many `Pets` (has many) and `Toys` (has many - polymorphic)
// He works in a Company (belongs to), he has a Manager (belongs to - single-table), and also managed a Team (has many - single-table)
// He speaks many languages (many to many) and has many friends (many to many - single-table)
// His pet also has one Toy (has one - polymorphic)
type User struct {
	gorm.Model
	Name      string
	Age       uint
	Birthday  *time.Time
	Account   Account
	Pets      []*Pet
	Toys      []Toy `gorm:"polymorphic:Owner"`
	CompanyID *int
	Company   Company
	ManagerID *uint
	Manager   *User
	Team      []User     `gorm:"foreignkey:ManagerID"`
	Languages []Language `gorm:"many2many:UserSpeak"`
	Friends   []*User    `gorm:"many2many:user_friends"`
	Active    bool
}

type Account struct {
	gorm.Model
	UserID sql.NullInt64
	Number string
}

type Pet struct {
	gorm.Model
	UserID *uint
	Name   string
	Toy    Toy `gorm:"polymorphic:Owner;"`
}

type Toy struct {
	gorm.Model
	Name      string
	OwnerID   string
	OwnerType string
}

type Company struct {
	ID   int
	Name string
}

type Language struct {
	Code string `gorm:"primarykey"`
	Name string
}

type GameUser struct {
	gorm.Model
	// 游戏昵称
	Nick string `gorm:"uniqueIndex;size:255"`
	// 联队
	Clan string `gorm:"size:255"`
	// 联队地址
	ClanUrl string `gorm:"size:255"`
	// 是否被封禁
	Banned *bool
	// 注册日期
	RegisterDate time.Time
	// 称号
	Title string `gorm:"size:255"`
	// 游戏等级
	Level  int
	StatAb UserStat `gorm:"embedded;embeddedPrefix:stat_ab_"`
	StatRb UserStat `gorm:"embedded;embeddedPrefix:stat_rb_"`
	StatSb UserStat `gorm:"embedded;embeddedPrefix:stat_sb_"`

	GroundRateAb GroundRate `gorm:"embedded;embeddedPrefix:rate_ground_ab_"`
	GroundRateRb GroundRate `gorm:"embedded;embeddedPrefix:rate_ground_rb_"`
	GroundRateSb GroundRate `gorm:"embedded;embeddedPrefix:rate_ground_sb_"`

	AviationRateAb AviationRate `gorm:"embedded;embeddedPrefix:rate_aviation_ab_"`
	AviationRateRb AviationRate `gorm:"embedded;embeddedPrefix:rate_aviation_rb_"`
	AviationRateSb AviationRate `gorm:"embedded;embeddedPrefix:rate_aviation_sb_"`

	FleetRateAb FleetRate `gorm:"embedded;embeddedPrefix:rate_fleet_ab_"`
	FleetRateRb FleetRate `gorm:"embedded;embeddedPrefix:rate_fleet_rb_"`
	FleetRateSb FleetRate `gorm:"embedded;embeddedPrefix:rate_fleet_sb_"`

	// TS街机效率值
	TsABRate float64
	// TS历史效率值
	TsRBRate float64
	// TS全真效率值
	TsSBRate float64
	// 安东星街机效率值
	AsABRate float64
	// 安东星历史效率值
	AsRBRate float64
	// 安东星全真效率值
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
