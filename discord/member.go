package discord

import "time"

type Member struct {
	User                       *User        `json:"user,omitempty"`
	Nick                       *string      `json:"nick,omitempty"`
	Avatar                     *string      `json:"avatar,omitempty"`
	Roles                      []Snowflake  `json:"roles"`
	JoinedAt                   time.Time    `json:"joined_at"`
	PremiumSince               *time.Time   `json:"premium_since,omitempty"`
	Deaf                       bool         `json:"deaf"`
	Mute                       bool         `json:"mute"`
	Pending                    *bool        `json:"pending,omitempty"`
	Permissions                *Permissions `json:"permissions,omitempty"`
	CommunicationDisabledUntil *time.Time   `json:"communication_disabled_until,omitempty"`
}
