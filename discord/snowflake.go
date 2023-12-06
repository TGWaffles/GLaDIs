package discord

import (
	"encoding/json"
	"strconv"
)

type Snowflake uint64

func GetSnowflake(id interface{}) (Snowflake, error) {
	switch id.(type) {
	case string:
		res, err := strconv.ParseInt(id.(string), 10, 64)
		return Snowflake(res), err
	case int:
		return Snowflake(id.(int)), nil
	case int64:
		return Snowflake(id.(int64)), nil
	default:
		return Snowflake(id.(uint64)), nil
	}
}

func (i *Snowflake) MarshalJSON() ([]byte, error) {
	return json.Marshal(strconv.FormatUint(uint64(*i), 10))
}

func (i *Snowflake) UnmarshalJSON(b []byte) error {
	// Try string first
	var s string
	if err := json.Unmarshal(b, &s); err == nil {
		value, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return err
		}
		*i = Snowflake(value)
		return nil
	}

	// Fallback to number
	return json.Unmarshal(b, (*uint64)(i))
}

func (i *Snowflake) String() string {
	return strconv.FormatUint(uint64(*i), 10)
}

func (i *Snowflake) Parse(s string) error {
	value, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return err
	}
	*i = Snowflake(value)
	return nil
}
