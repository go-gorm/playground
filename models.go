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


type Model struct {
	ID        uint       `gorm:"primary_key," json:"id"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"type:datetime"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"type:datetime"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"  gorm:"type:datetime"`
}

type AdSet struct {
	AdsetId          uint64     `json:"adset_id" gorm:"type:bigint;primary_key;autoIncrement:false"`
	Source           string     `json:"source" gorm:"type:varchar(20);primary_key;autoIncrement:false"`
	AccountCurrency  string     `json:"account_currency" gorm:"type:varchar(20)"`
	AccountId        uint64     `json:"account_id" gorm:"type:bigint"`
	CampaignId       uint64     `json:"campaign_id" gorm:"type:bigint"`
	AdsetName        string     `json:"adset_name"`
	BuyingType       string     `json:"buying_type"`
	Objective        string     `json:"objective"`
	CurrentBid       float64    `json:"current_bid"`
	AdsetStatus      string     `json:"adset_status"`
	DailyBudget      float64    `json:"daily_budget"`
	BidStrategy      string     `json:"bid_strategy"`
	CBO              bool       `json:"cbo"`
	Region           string     `json:"region" gorm:"type:json"`
	Country          string     `json:"country" gorm:"type:varchar(20)"`
	DomainId         *uint16    `json:"domain_id" gorm:"type:smallint(5)"`
	RedirectDomainId *uint16    `json:"redirect_domain_id" gorm:"type:smallint(5)"`
	Device           string     `json:"device" gorm:"type:json"`
	Extra            string     `json:"extra" gorm:"type:json"`
	StartDate        *time.Time `json:"start_date"`
	CreatedAt        *time.Time `json:"created_at,omitempty" gorm:"type:datetime"`
	UpdatedAt        *time.Time `json:"updated_at,omitempty" gorm:"type:datetime"`
}

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
