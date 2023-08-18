package main

type Account struct {
	Network Network `gorm:"embedded;embeddedPrefix:network_"`
	Peers   []Peer  `json:"-" gorm:"foreignKey:AccountID;references:ID"`
	ID      string  `gorm:"primaryKey"`
}

type Network struct {
	ID string
}

type Peer struct {
	ID        string `gorm:"primaryKey"`
	AccountID string
}
