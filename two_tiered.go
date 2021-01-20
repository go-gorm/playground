package main

import (
	"fmt"
	"sort"

	"gorm.io/gorm"
)

// TwoTieredEntry represents an entry.
type TwoTieredEntry struct {
	EntryPK uint64           `gorm:"primary_key;autoIncrement:false;column:entry_pk;type:INT8;"`
	Links   map[string]int64 `gorm:"-"`
}

// TwoTieredEntryDb holds one row from the entries table.
type TwoTieredEntryDb struct {
	TwoTieredEntry
	LinksDb  []TwoTieredLinkDb `gorm:"foreignKey:EntryPK;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// TableName overrides the table name used by TwoTieredEntryDb to `entries`.
func (TwoTieredEntryDb) TableName() string {
	return "two_entries"
}

func (edb *TwoTieredEntryDb) AfterFind(tx *gorm.DB) (err error) {
	if tx.Error == nil {
		edb.Links = make(map[string]int64)
		for _, l := range edb.LinksDb {
			if l.Key == "" {
				return fmt.Errorf("empty key for entry links")
			}
			edb.Links[l.Key] = l.Link
		}
		edb.LinksDb = nil
	}
	return
}

func (edb *TwoTieredEntryDb) BeforeSave(_ *gorm.DB) (err error) {
	var sorted []string
	for key := range edb.Links {
		sorted = append(sorted, key)
	}
	sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })
	for _, key := range sorted {
		edb.LinksDb = append(edb.LinksDb, TwoTieredLinkDb{EntryPK: edb.EntryPK, Key: key, Link: edb.Links[key]})
	}
	return
}

func (edb *TwoTieredEntryDb) AfterSave(_ *gorm.DB) (err error) {
	edb.LinksDb = nil
	return
}

// TwoTieredLinkDb holds one row from the links table.
type TwoTieredLinkDb struct {
	EntryPK uint64 `gorm:"primary_key;autoIncrement:false;column:entry_pk;type:INT8;"`
	Key     string `gorm:"primary_key;autoIncrement:false;column:key;type:VARCHAR;size:64;"`
	Link    int64  `gorm:"column:link;type:INT4;"`
}

// TableName overrides the table name used by LinkDb to `links`.
func (TwoTieredLinkDb) TableName() string {
	return "two_links"
}
