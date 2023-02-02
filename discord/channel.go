package discord

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

type ChannelType uint8
