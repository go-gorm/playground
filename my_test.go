package main

import (
	"fmt"
	"testing"

	gomysql "github.com/go-sql-driver/mysql"
)

func getDatabaseSourceName() string {
	username := "gorm"
	password := "gorm"
	host := "127.0.0.1"
	port := "9910"
	name := "gorm"

    cfg := gomysql.Config{
        User:   username,
        Passwd: password,
        Net:    "tcp",
        Addr:   host + ":" + port,
        DBName: name,
        Params: map[string]string{
            "charset": "utf8mb4",
            "parseTime": "True",
            "loc": "Local",
            "multiStatements": "true",
        },
    }

    url := cfg.FormatDSN()

	return url
}

type Organisation2 struct {
    ID    string `gorm:"primaryKey"`
	myName  string `gorm:"column:NAME"` // <<<< THIS SHOULD WORK, BUT ISN'T
}

func (Organisation2) TableName() string {
    return "T_ORGANISATION"
}

func TestNameLikeOtherTests(t *testing.T) {

	// create table and populate it
	DB.Exec("CREATE TABLE IF NOT EXISTS T_ORGANISATION (ID VARCHAR(36) NOT NULL, NAME VARCHAR(100) NOT NULL, PRIMARY KEY (ID) ) ")
	fmt.Println("created table")
	DB.Exec("DELETE FROM T_ORGANISATION")
	fmt.Println("deleted all organisations")
	DB.Exec("INSERT INTO T_ORGANISATION (ID, NAME) VALUES ('A', 'NAME') ")
	fmt.Println("inserted one organisation")
	
	var result Organisation2
	if err := DB.First(&result).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if len(result.myName) <= 0 {
		t.Errorf("Expected non-empty, got '%q'", result.myName) 
	}



}


