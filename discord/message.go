package discord

import "time"

const (
	DefaultMessageType MessageType = iota
	RecipientAddMessageType
	RecipientRemoveMessageType
	CallMessageType
	ChannelNameChangeMessageType
	ChannelIconChangeMessageType
	ChannelPinnedMessageMessageType
	GuildMemberJoinMessageType
	GuildBoostMessageType
	GuildBoostTier1MessageType
	GuildBoostTier2MessageType
	GuildBoostTier3MessageType
	ChannelFollowAddMessageType
	GuildDiscoveryDisqualifiedMessageType = iota + 1
	GuildDiscoveryReQualifiedMessageType
	GuildDiscoveryGracePeriodInitialWarningMessageType
	GuildDiscoveryGracePeriodFinalWarningMessageType
	ThreadCreatedMessageType
	ReplyMessageType
	ChatInputCommandMessageType
	ThreadStarterMessageMessageType
	GuildInviteReminderMessageType
	ContextMenuCommandMessageType
	AutoModerationActionMessageType
	RoleSubscriptionPurchaseMessageType
	InteractionPremiumUpsellMessageType
	GuildApplicationPremiumSubscriptionMessageType = iota + 6
)

const (
	JoinMessageActivityType MessageActivityType = iota + 1
	SpectateMessageActivityType
	ListenMessageActivityType
	JoinRequestMessageActivityType
)

type Message struct {
	Id              Snowflake        `json:"id"`
	ChannelId       Snowflake        `json:"channel_id"`
	Author          *User            `json:"author"`
	Content         string           `json:"content"`
	Timestamp       time.Time        `json:"timestamp"`
	EditedTimestamp time.Time        `json:"edited_timestamp"`
	Tts             bool             `json:"tts"`
	MentionEveryone bool             `json:"mention_everyone"`
	Mentions        []User           `json:"mentions"`
	MentionRoles    []Snowflake      `json:"mention_roles"`
	MentionChannels []ChannelMention `json:"mention_channels"`
	Attachments     []Attachment     `json:"attachments"`
	Embeds          []Embed          `json:"embeds"`
	Reactions       *[]Reaction      `json:"reactions,omitempty"`
	Nonce           interface{}      `json:"nonce,omitempty"`
	Pinned          bool             `json:"pinned"`
	WebhookId       *Snowflake       `json:"webhook_id,omitempty"`
	Type            MessageType      `json:"type"`
}

type ChannelMention struct {
	Id      Snowflake   `json:"id"`
	GuildId Snowflake   `json:"guild_id"`
	Type    ChannelType `json:"type"`
	Name    string      `json:"name"`
}

type MessageActivity struct {
	Type    MessageActivityType `json:"type"`
	PartyId string              `json:"party_id"`
}

type MessageActivityType int

type MessageType uint8
