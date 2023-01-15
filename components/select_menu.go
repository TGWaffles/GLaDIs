package components

import "github.com/tgwaffles/lambda-discord-interactions-go/discord"

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

type SelectMenu struct {
	Type         ComponentType   `json:"type"`
	CustomId     string          `json:"custom_id"`
	Options      *[]SelectOption `json:"options,omitempty"`
	ChannelTypes *[]ChannelType  `json:"channel_types,omitempty"`
	Placeholder  *string         `json:"placeholder,omitempty"`
	MinValues    *uint8          `json:"min_values,omitempty"`
	MaxValues    *uint8          `json:"max_values,omitempty"`
	Disabled     *bool           `json:"disabled,omitempty"`
}

type SelectOption struct {
	Label       string         `json:"label"`
	Value       string         `json:"value"`
	Description *string        `json:"description,omitempty"`
	Emoji       *discord.Emoji `json:"emoji,omitempty"`
	Default     *bool          `json:"default,omitempty"`
}

type ChannelType uint8
