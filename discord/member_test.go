package discord

import (
	"encoding/json"
	"testing"
	"time"
)

func TestRoleSnowflakeParsing(t *testing.T) {
	member := Member{}
	err := json.Unmarshal([]byte(`{"roles": ["1234"]}`), &member)
	if err != nil {
		t.Error(err)
	}
	if member.Roles[0] != 1234 {
		t.Errorf("Expected member.Roles[0] to be 1234, got %d", member.Roles[0])
	}
}

func TestDateParsing(t *testing.T) {
	member := Member{}
	err := json.Unmarshal([]byte(`{"joined_at": "2015-04-26T06:26:56.936000+00:00"}`), &member)
	if err != nil {
		t.Error(err)
	}

	expected := time.Date(2015, 4, 26, 6, 26, 56, 936000000, time.UTC)
	if member.JoinedAt.UTC() != expected {
		t.Errorf("Expected member.JoinedAt to be %s, got %s", expected, member.JoinedAt)
	}
}
