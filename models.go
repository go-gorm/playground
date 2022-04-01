package main

type User struct {
	ID       uint64 `faker:"-"`
	Username string
}

type Game struct {
	ID   uint64 `faker:"-"`
	Name string
}

type Set struct {
	ID         uint64 `faker:"-"`
	GameID     uint64 `faker:"-"`
	Name       string
	ParentCode string `gorm:"default:(-)"`
	Game       *Game
}

type Article struct {
	ID    uint64 `faker:"-"`
	SetID uint64 `faker:"-"`
	Name  string
	Set   *Set
}

type Product struct {
	ID        uint64 `faker:"-"`
	UserID    uint64 `faker:"-"`
	ArticleID uint64 `faker:"-"`
	Code      string
	User      *User
	Article   *Article
}
