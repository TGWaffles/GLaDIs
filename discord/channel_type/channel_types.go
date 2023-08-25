package channel_type

const (
	GuildText ChannelType = iota
	DM
	GuildVoice
	GroupDM
	GuildCategory
	GuildAnnouncement
	AnnouncementThread = iota + 4
	PublicThread
	PrivateThread
	GuildStageVoice
	GuildDirectory
	GuildForum
)

type ChannelType uint8
