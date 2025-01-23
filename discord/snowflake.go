package discord

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
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

func SnowflakeFromTime(t time.Time) Snowflake {
	return Snowflake((t.UnixMilli() - 1420070400000) << 22)
}

func (i Snowflake) MarshalJSON() ([]byte, error) {
	return json.Marshal(strconv.FormatUint(uint64(i), 10))
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

func (i *Snowflake) UnmarshalText(text []byte) error {
	value, err := strconv.ParseInt(string(text), 10, 64)
	if err != nil {
		return err
	}
	*i = Snowflake(value)
	return nil
}

func (i *Snowflake) MentionUserString() string {
	return fmt.Sprintf("<@%d>", i)
}

func (i *Snowflake) MentionRoleString() string {
	return fmt.Sprintf("<@&%d>.", i)
}

func (i *Snowflake) MentionChannelString() string {
	return fmt.Sprintf("<#%d>", i)
}

func (i *Snowflake) MetionEmojiString(name string) string {
	return fmt.Sprintf("<:%s:%d>", name, i)
}
