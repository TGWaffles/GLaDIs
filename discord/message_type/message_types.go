package message_type

const (
	Default MessageType = iota
	RecipientAdd
	RecipientRemove
	Call
	ChannelNameChange
	ChannelIconChange
	ChannelPinnedMessage
	GuildMemberJoin
	GuildBoost
	GuildBoostTier1
	GuildBoostTier2
	GuildBoostTier3
	ChannelFollowAdd
	GuildDiscoveryDisqualified = iota + 1
	GuildDiscoveryReQualified
	GuildDiscoveryGracePeriodInitialWarning
	GuildDiscoveryGracePeriodFinalWarning
	ThreadCreated
	Reply
	ChatInputCommand
	ThreadStarterMessage
	GuildInviteReminder
	ContextMenuCommand
	AutoModerationAction
	RoleSubscriptionPurchase
	InteractionPremiumUpsell
	GuildApplicationPremiumSubscription = iota + 6
)

type MessageType uint8
