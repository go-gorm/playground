package main

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cast"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type TinyInt uint8

// GormDataType returns gorm common data type. This type is used for the field's column type.
func (TinyInt) GormDataType() string {
	return string(schema.Uint)
}

// GormDBDataType returns gorm DB data type based on the current using database.
func (TinyInt) GormDBDataType(db *gorm.DB, field *schema.Field) (dataType string) {
	switch db.Dialector.Name() {
	case "sqlite":
		dataType = "integer"
	case "mysql":
		dataType = "tinyint unsigned"
	case "postgres", "dm":
		dataType = "smallint"
	case "sqlserver":
		dataType = "tinyint"
	case "oracle":
		dataType = "NUMBER(3, 0)"
	default:
		dataType = getGormTypeFromTag(field)
	}
	return
}

func getGormTypeFromTag(field *schema.Field) (dataType string) {
	if field != nil {
		if val, ok := field.TagSettings["TYPE"]; ok {
			dataType = strings.ToLower(val)
		}
	}
	return
}

// Scan implements sql.Scanner interface and scans value into TinyInt.
func (tinyint *TinyInt) Scan(src interface{}) error {
	if src == nil {
		*tinyint = 0
		return nil
	}

	v, err := cast.ToUint8E(src)
	if err != nil {
		return fmt.Errorf("failed to parse TinyInt value: %v", src)
	}
	*tinyint = TinyInt(v)
	return nil
}

// Value implements driver.Valuer interface and returns uint8 format of TinyInt.
func (tinyint TinyInt) Value() (driver.Value, error) {
	if tinyint == 0 {
		return nil, nil
	}

	return int64(tinyint), nil
}

// String implements fmt.Stringer interface.
func (tinyint TinyInt) String() string {
	return strconv.FormatUint(uint64(tinyint), 10)
}

//goland:noinspection GoMixedReceiverTypes
func (tinyint TinyInt) Int() int {
	return int(tinyint)
}

// MarshalJSON implements json.Marshaler to convert TinyInt to json serialization.
func (tinyint TinyInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(tinyint.Int())
}

// UnmarshalJSON implements json.Unmarshaler to deserialize json data.
func (tinyint *TinyInt) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}
	s := string(data)
	if value, err := strconv.Unquote(s); err == nil {
		return tinyint.Scan(value)
	}
	return tinyint.Scan(s)
}
