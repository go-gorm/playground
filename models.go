package main

type Folder struct {
	ID       int
	Name     string `gorm:"non null"`
	ParentID *int
	Parent   *Folder `gorm:"foreignkey:ParentID"`
}
