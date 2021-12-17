package main

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: postgres

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"net"
	"testing"
)

type Inet struct {
	net.IP
}

func ParseIP(s string) (ip Inet, err error) {
	i := net.ParseIP(s)
	if i == nil {
		return ip, fmt.Errorf("can't parse IP %s", s)
	}
	ip.IP = i
	return ip, err
}
func MustParseIP(s string) (ip Inet) {
	ip, err := ParseIP(s)
	if err != nil {
		panic(err)
	}
	return ip
}

func (i *Inet) Scan(value interface{}) error {
	i.IP = nil
	if value == nil {
		return nil
	}
	ipAsBytes, ok := value.([]byte)
	if !ok {
		return errors.New("Could not convert scanned value to bytes")
	}
	parsedIP := net.ParseIP(string(ipAsBytes))
	if parsedIP == nil {
		return nil
	}
	i.IP = parsedIP
	return nil
}
func (i Inet) Value() (driver.Value, error) {
	if i.IP == nil {
		return nil, nil
	}
	return []byte(i.IP.String()), nil
}

type AddrTest struct {
	IP Inet `gorm:"type:inet" json:"ip"`
}

func TestCustomType(t *testing.T) {
	err := DB.AutoMigrate(AddrTest{})
	if err != nil {
		t.Fatalf("Failed to create table: %s", err)
	}
	DB.Create(AddrTest{IP: MustParseIP("1.2.3.4")})
}
