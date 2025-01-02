package types

import (
	"database/sql/driver"
	"encoding/json"
)

type StringArray []string

func (a StringArray) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *StringArray) Scan(value interface{}) error {
	if value == nil {
		*a = StringArray{}
		return nil
	}
	return json.Unmarshal(value.([]byte), a)
}
