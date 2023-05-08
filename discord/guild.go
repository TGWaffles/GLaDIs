package discord

type Guild struct {
	Id                          Snowflake      `json:"id"`
	Name                        string         `json:"name"`
	Icon                        *string        `json:"icon,omitempty"`
	IconHash                    *string        `json:"icon_hash,omitempty"`
	Splash                      *string        `json:"splash,omitempty"`
	DiscoverySplash             *string        `json:"discovery_splash,omitempty"`
	Owner                       *bool          `json:"owner,omitempty"`
	OwnerId                     Snowflake      `json:"owner_id"`
	Permissions                 *Permissions   `json:"permissions,omitempty"`
	Region                      *string        `json:"region,omitempty"`
	AfkChannelId                *Snowflake     `json:"afk_channel_id,omitempty"`
	AfkTimeout                  int            `json:"afk_timeout"`
	WidgetEnabled               *bool          `json:"widget_enabled,omitempty"`
	WidgetChannelId             *Snowflake     `json:"widget_channel_id,omitempty"`
	VerificationLevel           int            `json:"verification_level"`
	DefaultMessageNotifications int            `json:"default_message_notifications"`
	ExplicitContentFilter       int            `json:"explicit_content_filter"`
	Roles                       []Role         `json:"roles"`
	Emojis                      []Emoji        `json:"emojis"`
	Features                    []string       `json:"features"`
	MfaLevel                    int            `json:"mfa_level"`
	ApplicationId               *Snowflake     `json:"application_id,omitempty"`
	SystemChannelId             *Snowflake     `json:"system_channel_id,omitempty"`
	SystemChannelFlags          int            `json:"system_channel_flags"`
	RulesChannelId              *Snowflake     `json:"rules_channel_id,omitempty"`
	MaxPresences                *int           `json:"max_presences,omitempty"`
	MaxMembers                  int            `json:"max_members"`
	VanityURLCode               *string        `json:"vanity_url_code,omitempty"`
	Description                 *string        `json:"description,omitempty"`
	Banner                      *string        `json:"banner,omitempty"`
	PremiumTier                 int            `json:"premium_tier"`
	PremiumSubscriptionCount    *int           `json:"premium_subscription_count,omitempty"`
	PreferredLocale             string         `json:"preferred_locale"`
	PublicUpdatesChannelId      *Snowflake     `json:"public_updates_channel_id,omitempty"`
	MaxVideoChannelUsers        *int           `json:"max_video_channel_users,omitempty"`
	MaxStageVideoChannelUsers   *int           `json:"max_stage_video_channel_users,omitempty"`
	ApproximateMemberCount      *int           `json:"approximate_member_count,omitempty"`
	ApproximatePresenceCount    *int           `json:"approximate_presence_count,omitempty"`
	WelcomeScreen               *WelcomeScreen `json:"welcome_screen,omitempty"`
	NsfwLevel                   int            `json:"nsfw_level"`
	Stickers                    []Sticker      `json:"stickers"`
	PremiumProgressBarEnabled   *bool          `json:"premium_progress_bar_enabled,omitempty"`
}

type WelcomeScreen struct {
	Description     *string                `json:"description,omitempty"`
	WelcomeChannels []WelcomeScreenChannel `json:"welcome_channels"`
}

type WelcomeScreenChannel struct {
	ChannelId   Snowflake  `json:"channel_id"`
	Description string     `json:"description"`
	EmojiId     *Snowflake `json:"emoji_id,omitempty"`
	EmojiName   *string    `json:"emoji_name,omitempty"`
}

type Sticker struct {
	Id          Snowflake  `json:"id"`
	PackId      *Snowflake `json:"pack_id,omitempty"`
	Name        string     `json:"name"`
	Description *string    `json:"description,omitempty"`
	Tags        *string    `json:"tags,omitempty"`
	Asset       *string    `json:"asset,omitempty"`
	Type        int        `json:"type"`
	FormatType  int        `json:"format_type"`
	Available   *bool      `json:"available,omitempty"`
	GuildId     *Snowflake `json:"guild_id,omitempty"`
	User        *User      `json:"user,omitempty"`
	SortValue   *int       `json:"sort_value,omitempty"`
}
