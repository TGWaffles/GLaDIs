package discord

type User struct {
	Id            Snowflake `json:"id"`
	Username      string    `json:"username"`
	Discriminator string    `json:"discriminator"`
	Avatar        *string   `json:"avatar,omitempty"`
	Bot           *bool     `json:"bot,omitempty"`
	System        *bool     `json:"system,omitempty"`
	MfaEnabled    *bool     `json:"mfa_enabled,omitempty"`
	Banner        *string   `json:"banner,omitempty"`
	AccentColor   *uint32   `json:"accent_color,omitempty"`
	Locale        *string   `json:"locale,omitempty"`
	Verified      *bool     `json:"verified,omitempty"`
	Email         *string   `json:"email,omitempty"`
	Flags         *uint32   `json:"flags,omitempty"`
	PremiumType   *uint8    `json:"premium_type,omitempty"`
	PublicFlags   *uint32   `json:"public_flags,omitempty"`
}
