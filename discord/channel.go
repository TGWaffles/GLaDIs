package discord

import "time"

const (
	GuildTextChannelType ChannelType = iota
	DMChannelType
	GuildVoiceChannelType
	GroupDMChannelType
	GuildCategoryChannelType
	GuildAnnouncementChannelType
	AnnouncementThreadChannelType = iota + 4
	PublicThreadChannelType
	PrivateThreadChannelType
	GuildStageVoiceChannelType
	GuildDirectoryChannelType
	GuildForumChannelType
)

const (
	RoleOverwriteType OverwriteType = iota
	MemberOverwriteType
)

type ChannelType uint8
type OverwriteType uint8

type Channel struct {
	Id                            Snowflake        `json:"id"`
	Type                          ChannelType      `json:"type"`
	GuildId                       *Snowflake       `json:"guild_id,omitempty"`
	Position                      *int             `json:"position,omitempty"`
	PermissionOverwrites          *[]Overwrite     `json:"permission_overwrites,omitempty"`
	Name                          *string          `json:"name,omitempty"`
	Topic                         *string          `json:"topic,omitempty"`
	Nsfw                          *bool            `json:"nsfw,omitempty"`
	LastMessageId                 *Snowflake       `json:"last_message_id,omitempty"`
	Bitrate                       *int             `json:"bitrate,omitempty"`
	UserLimit                     *int             `json:"user_limit,omitempty"`
	RateLimitPerUser              *int             `json:"rate_limit_per_user,omitempty"`
	Recipients                    *[]User          `json:"recipients,omitempty"`
	Icon                          *string          `json:"icon,omitempty"`
	OwnerId                       *Snowflake       `json:"owner_id,omitempty"`
	ApplicationId                 *Snowflake       `json:"application_id,omitempty"`
	Managed                       *bool            `json:"managed,omitempty"`
	ParentId                      *Snowflake       `json:"parent_id,omitempty"`
	LastPinTimestamp              *time.Time       `json:"last_pin_timestamp,omitempty"`
	RtcRegion                     *string          `json:"rtc_region,omitempty"`
	VideoQualityMode              *int             `json:"video_quality_mode,omitempty"`
	MessageCount                  *int             `json:"message_count,omitempty"`
	MemberCount                   *int             `json:"member_count,omitempty"`
	ThreadMetadata                *ThreadMetadata  `json:"thread_metadata,omitempty"`
	Member                        *ThreadMember    `json:"member,omitempty"`
	DefaultAutoArchiveDuration    *int             `json:"default_auto_archive_duration,omitempty"`
	Permissions                   *string          `json:"permissions,omitempty"`
	Flags                         *int             `json:"flags,omitempty"`
	TotalMessageSent              *int             `json:"total_message_sent,omitempty"`
	AvailableTags                 *[]Tag           `json:"available_tags,omitempty"`
	AppliedTags                   *[]Snowflake     `json:"applied_tags,omitempty"`
	DefaultReactionEmoji          *DefaultReaction `json:"default_reaction_emoji,omitempty"`
	DefaultThreadRateLimitPerUser *int             `json:"default_thread_rate_limit_per_user,omitempty"`
	DefaultSortOrder              *int             `json:"default_sort_order,omitempty"`
	DefaultForumLayout            *int             `json:"default_forum_layout,omitempty"`
}

type Overwrite struct {
	Id    Snowflake     `json:"id"`
	Type  OverwriteType `json:"type"`
	Allow string        `json:"allow"`
	Deny  string        `json:"deny"`
}

type ThreadMetadata struct {
	Archived            bool       `json:"archived"`
	AutoArchiveDuration int        `json:"auto_archive_duration"`
	ArchiveTimestamp    time.Time  `json:"archive_timestamp"`
	Locked              bool       `json:"locked"`
	Invitable           *bool      `json:"invitable,omitempty"`
	CreateTimestamp     *time.Time `json:"create_timestamp,omitempty"`
}

type ThreadMember struct {
	Id            *Snowflake `json:"id,omitempty"`
	UserId        *Snowflake `json:"user_id,omitempty"`
	JoinTimestamp time.Time  `json:"join_timestamp"`
	Flags         int        `json:"flags"`
	Member        *Member    `json:"member,omitempty"`
}

type Tag struct {
	Id        Snowflake  `json:"id"`
	Name      string     `json:"name"`
	Moderated bool       `json:"moderated"`
	EmojiId   *Snowflake `json:"emoji_id,omitempty"`
	EmojiName *string    `json:"emoji_name,omitempty"`
}

type DefaultReaction struct {
	EmojiId   *Snowflake `json:"emoji_id,omitempty"`
	EmojiName *string    `json:"emoji_name,omitempty"`
}
