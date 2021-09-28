package identifier

import (
	"database/sql/driver"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/oklog/ulid/v2"
)

type ID ulid.ULID

var (
	Nil ID
)

//GetNewULIDString get Lowercase ulid
func IsNil(id ID) bool {
	return id == Nil
}

//GetNewULIDString get Lowercase ulid
func GetNewIDString() string {
	return GetNewID().String()
}

//GetNewULIDString get Lowercase ulid
func GetNewID() ID {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return ID(id)
}

//FromString ulid wrapper method for converting to ulid
func FromString(s string) (id ID, err error) {
	ulid, err := ulid.ParseStrict(s)
	if err != nil {
		log.Printf("ERR: failed to ulid method: FromString :%v", err)
		return
	}

	id = ID(ulid)

	return
}

//String convert to string
func (id ID) String() string {
	return strings.ToLower(ulid.ULID(id).String())
}

// MarshalJSON -> convert to json string
func (id ID) MarshalJSON() ([]byte, error) {
	str := "\"" + id.String() + "\""
	return []byte(str), nil
}

// UnmarshalJSON -> convert from json string
func (id *ID) UnmarshalJSON(by []byte) error {
	s := string(by)
	s = strings.Replace(s, "\"", "", 2)
	ulid, err := ulid.ParseStrict(s)
	*id = ID(ulid)
	return err
}

//Gorm Interface
// GormDataType -> sql data type for gorm
//TODO ID?
func (ID) GormDataType() string {
	return "binary(16)"
}

// Scan -> scan value into BinaryID
func (id *ID) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal JSONB, method: identifier.Scan, value: %v", value)
	}

	var ulid ulid.ULID
	err := ulid.UnmarshalBinary(bytes)

	*id = ID(ulid)
	return err
}

// Value -> return BinaryID to []bytes binary(16)
func (id ID) Value() (driver.Value, error) {
	return ulid.ULID(id).MarshalBinary()
}
