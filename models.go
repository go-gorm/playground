package main

type Account struct {
	ID      string  `gorm:"primaryKey"`
	Network Network `gorm:"embedded;embeddedPrefix:network_"`
	Peers   []Peer  `json:"-" gorm:"foreignKey:AccountID;references:ID"`
}

type Network struct {
	ID string
}

type Peer struct {
	ID        string `gorm:"primaryKey"`
	AccountID string
}
