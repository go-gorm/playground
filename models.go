package main

import (
	"crypto/md5"
	"database/sql"
	"database/sql/driver"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"github.com/pkg/errors"
	"gorm.io/gorm/schema"
	"time"

	"gorm.io/gorm"
)

type CryptoFieldJsonStorage struct {
	Plaintext *string `json:"plaintext"`
	FullHash  *string `json:"full_hash"`
}

type Crypto struct {
	data        []byte
	jsonStorage *CryptoFieldJsonStorage
}

func (c *Crypto) GormDataType() string {
	return "json"
}

// GormDBDataType gorm db data type
func (c *Crypto) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite":
		return "JSON"
	case "mysql":
		return "JSON"
	case "postgres":
		return "JSONB"
	}
	return ""
}

// Implements driver.Scanner
func (c *Crypto) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("invalid data type")
	}

	c.jsonStorage = &CryptoFieldJsonStorage{}
	err := json.Unmarshal(source, c.jsonStorage)
	if err != nil {
		return err
	}

	//if c.jsonStorage.Plaintext != nil {
	//	c.data, err = base64.StdEncoding.DecodeString(*c.jsonStorage.Plaintext)
	//	if err != nil {
	//		return err
	//	}
	//}
	return nil
}

// Implements driver.Valuer
func (c Crypto) Value() (driver.Value, error) {
	if c.data == nil {
		return nil, nil // THIS
	}
	plaintext := base64.StdEncoding.EncodeToString(c.data)
	c.jsonStorage.Plaintext = &plaintext

	md5Hash := md5.Sum(c.data)
	fullHash := hex.EncodeToString(md5Hash[:])
	c.jsonStorage.FullHash = &fullHash

	jsonStorage, err := json.Marshal(c.jsonStorage)
	if err != nil {
		return nil, err
	}
	return jsonStorage, nil
}

func NewCrypto(data []byte) Crypto {
	return Crypto{
		data:        data,
		jsonStorage: &CryptoFieldJsonStorage{},
	}
}


// User has one `Account` (has one), many `Pets` (has many) and `Toys` (has many - polymorphic)
// He works in a Company (belongs to), he has a Manager (belongs to - single-table), and also managed a Team (has many - single-table)
// He speaks many languages (many to many) and has many friends (many to many - single-table)
// His pet also has one Toy (has one - polymorphic)
type User struct {
	gorm.Model
	Name      Crypto
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
