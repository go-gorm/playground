package main

import (
	"fmt"
	"sort"

	"gorm.io/gorm"
)

type ThreeTieredParent struct {
	ParentPK uint64              `gorm:"primary_key;autoIncrement:false;column:parent_pk;type:INT8;"`
	Entries  []*ThreeTieredEntry `gorm:"-"`
}

type ThreeTieredParentDb struct {
	ThreeTieredParent
	EntriesDb []*ThreeTieredEntryDb `gorm:"foreignKey:ParentPK;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// TableName overrides the table name used by ParentDb to `parents`.
func (ThreeTieredParentDb) TableName() string {
	return "three_parents"
}

func (tdb *ThreeTieredParentDb) AfterFind(tx *gorm.DB) (err error) {
	if tx.Error == nil {
		for _, edb := range tdb.EntriesDb {
			edb.ParentPK = tdb.ParentPK
			tdb.Entries = append(tdb.Entries, &edb.ThreeTieredEntry)
		}
		tdb.EntriesDb = nil
	}
	return
}

func (tdb *ThreeTieredParentDb) BeforeSave(_ *gorm.DB) (err error) {
	for _, e := range tdb.Entries {
		tdb.EntriesDb = append(tdb.EntriesDb, &ThreeTieredEntryDb{ThreeTieredEntry: *e, ParentPK: tdb.ParentPK})
	}
	return
}

func (tdb *ThreeTieredParentDb) AfterSave(_ *gorm.DB) (err error) {
	tdb.EntriesDb = nil
	return
}

// Entry represents an entry.
type ThreeTieredEntry struct {
	EntryPK uint64           `gorm:"primary_key;autoIncrement:false;column:entry_pk;type:INT8;"`
	Links   map[string]int64 `gorm:"-"`
}

// ThreeTieredEntryDb holds one row from the entries table.
type ThreeTieredEntryDb struct {
	ThreeTieredEntry
	ParentPK uint64              `gorm:"column:parent_pk;type:INT8;"`
	LinksDb  []ThreeTieredLinkDb `gorm:"foreignKey:EntryPK;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// TableName overrides the table name used by EntryDb to `entries`.
func (ThreeTieredEntryDb) TableName() string {
	return "three_entries"
}

func (edb *ThreeTieredEntryDb) AfterFind(tx *gorm.DB) (err error) {
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

func (edb *ThreeTieredEntryDb) BeforeSave(_ *gorm.DB) (err error) {
	var sorted []string
	for key := range edb.Links {
		sorted = append(sorted, key)
	}
	sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })
	for _, key := range sorted {
		edb.LinksDb = append(edb.LinksDb, ThreeTieredLinkDb{EntryPK: edb.EntryPK, Key: key, Link: edb.Links[key]})
	}
	return
}

func (edb *ThreeTieredEntryDb) AfterSave(_ *gorm.DB) (err error) {
	edb.LinksDb = nil
	return
}

// ThreeTieredLinkDb holds one row from the links table.
type ThreeTieredLinkDb struct {
	EntryPK uint64 `gorm:"primary_key;autoIncrement:false;column:entry_pk;type:INT8;"`
	Key     string `gorm:"primary_key;autoIncrement:false;column:key;type:VARCHAR;size:64;"`
	Link    int64  `gorm:"column:link;type:INT4;"`
}

// TableName overrides the table name used by ThreeTieredLinkDb to `links`.
func (ThreeTieredLinkDb) TableName() string {
	return "three_links"
}
