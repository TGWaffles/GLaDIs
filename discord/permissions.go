package discord

import (
	"encoding/json"
	"strconv"
)

type Permissions uint64

func (i *Permissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(strconv.FormatUint(uint64(*i), 10))
}

func (i *Permissions) UnmarshalJSON(b []byte) error {
	// Try string first
	var s string
	if err := json.Unmarshal(b, &s); err == nil {
		value, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return err
		}
		*i = Permissions(value)
		return nil
	}

	// Fallback to number
	return json.Unmarshal(b, (*uint64)(i))
}
