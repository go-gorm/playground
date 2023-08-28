package main

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"gorm.io/gorm/schema"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}

const encSuffix = "Encrypted"

type EncString struct {
	plain  string
	cipher string
}

func NewEncString(plain string) *EncString {
	return &EncString{plain: plain}
}

func (es *EncString) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue interface{}) (interface{}, error) {
	return es.plain + encSuffix, nil
	//return nil, nil
}

func (es *EncString) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue interface{}) error {
	switch value := dbValue.(type) {
	case []byte:
		es.cipher = string(value)
	case string:
		es.cipher = value
	default:
		fmt.Println("error")
		return fmt.Errorf("unsupported type nil")
	}
	es.plain = es.cipher
	if lastInd := strings.Index(es.cipher, encSuffix); lastInd != -1 {
		es.plain = es.cipher[:lastInd]
	}
	return nil
}

type Model1 struct {
	ID  int
	Enc *EncString `gorm:"type:varchar(128)"`
	//Contracts map[string]interface{} `gorm:"serializer:json"`
	CreatedAt time.Time
	UpdatedAt time.Time `gorm:"column:updated_at;not null"`
}

func TestEncString(t *testing.T) {
	db := DB
	//db.Migrator().DropTable(&Model1{})
	//db.Migrator().AutoMigrate(&Model1{})
	m1 := &Model1{Enc: NewEncString("peikai"), UpdatedAt: time.Now().Add(-time.Hour)}
	db.Create(m1)
	update := &Model1{Enc: NewEncString("peikaiUpdate"), UpdatedAt: time.Now().Add(-time.Hour)}
	check(db.Model(&Model1{}).Where(m1).UpdateColumns(update).Error)
	var m2 Model1
	m2.ID = 19
	check(db.Take(&m2).Error)
	fmt.Printf("%+v\n", m2)

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func TestUnmarshal(t *testing.T) {
	type model struct {
		Name string
		A    map[string]any
	}
	var m model
	fmt.Println(&m)
	err := json.Unmarshal([]byte(`{"Name":"peikai", "A":{"Name":"peikai", "A":678}}`), &m)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(m)
	}
}

func Test1(t *testing.T) {
	db := DB
	db.Migrator().DropTable(&Model1{})
	db.AutoMigrate(&Model1{})
	s1, s2 := "peikai1", "peikai2"
	db.Create(&Model1{Enc: NewEncString(s1)})
	db.Create(&Model1{Enc: NewEncString(s2)})

	var m1 Model1
	m1.ID = 1
	fieldSerializer := db.Find(&m1).Statement.Schema.LookUpField("enc").Serializer.(*EncString)
	require.Equal(t, m1.Enc.plain, s1)
	t.Logf("m1.Enc.plain = %s", m1.Enc.plain)
	N := 1000
	/*	go func() {
		var m1 Model1
		m1.ID = 1
		for i := 0; i < N; i++ {
			db.First(&m1)
			if m1.Enc == fieldSerializer {
				t.Errorf("m1.Enc == field.Serializer: %p", m1.Enc)

			}
			if m1.Enc.plain != s1 {
				t.Errorf("m1.Enc.plain != peikai1")
			}
		}
	}()*/
	go func() {
		var m2 Model1
		m2.ID = 2
		for i := 0; i < N; i++ {
			db.First(&m2)
			if m2.Enc == fieldSerializer {
				t.Errorf("m2.Enc == field.Serializer happened at iteration %d: %p", i, m2.Enc)
				t.Errorf("m1.Enc.plain = %s", m1.Enc.plain)
				require.Equal(t, m1.Enc.plain, s1)
			}
		}
	}()
	select {}
}
