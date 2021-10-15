package types

import (
	"database/sql/driver"
	"fmt"
)

//KeyValueMap will be translated into semicolon separated key=value
type KeyValueMap map[string]string

func (k *KeyValueMap) Scan(value interface{}) error {
	kvString, ok := value.(string)
	if !ok {
		return fmt.Errorf("%v is not a string", value)
	}

	*k = ToKeyValueMap(kvString)
	return nil
}

// Value return json value, implement driver.Valuer interface
func (k KeyValueMap) Value() (driver.Value, error) {
	return FromKeyValueMap(k)
}
