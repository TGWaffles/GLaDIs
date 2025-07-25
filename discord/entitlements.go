package discord

import "time"

type Entitlement struct {
	ID            Snowflake  `json:"id"`
	SkuID         Snowflake  `json:"sku_id"`
	ApplicationID Snowflake  `json:"application_id"`
	UserID        *Snowflake `json:"user_id,omitempty"`
	Type          int        `json:"type"`
	Deleted       bool       `json:"deleted"`
	StartsAt      *time.Time `json:"starts_at,omitempty"`
	EndsAt        *time.Time `json:"ends_at,omitempty"`
	GuildID       *Snowflake `json:"guild_id,omitempty"`
	Consumed      *bool      `json:"consumed,omitempty"`
}
