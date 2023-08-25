package discord

import (
	"github.com/tgwaffles/gladis/discord/channel_type"
	"github.com/tgwaffles/gladis/discord/message_activity_type"
	"github.com/tgwaffles/gladis/discord/message_type"
	"time"
)

type Message struct {
	Id                Snowflake                `json:"id"`
	ChannelId         Snowflake                `json:"channel_id"`
	Author            *User                    `json:"author"`
	Content           string                   `json:"content"`
	Timestamp         time.Time                `json:"timestamp"`
	EditedTimestamp   time.Time                `json:"edited_timestamp"`
	Tts               bool                     `json:"tts"`
	MentionEveryone   bool                     `json:"mention_everyone"`
	Mentions          []User                   `json:"mentions"`
	MentionRoles      []Snowflake              `json:"mention_roles"`
	MentionChannels   []ChannelMention         `json:"mention_channels"`
	Attachments       []Attachment             `json:"attachments"`
	Embeds            []Embed                  `json:"embeds"`
	Reactions         []Reaction               `json:"reactions,omitempty"`
	Nonce             interface{}              `json:"nonce,omitempty"`
	Pinned            bool                     `json:"pinned"`
	WebhookId         *Snowflake               `json:"webhook_id,omitempty"`
	Type              message_type.MessageType `json:"type"`
	Activity          *MessageActivity         `json:"activity,omitempty"`
	Application       *Application             `json:"application,omitempty"`
	ApplicationId     *Snowflake               `json:"application_id,omitempty"`
	MessageReference  *MessageReference        `json:"message_reference,omitempty"`
	Flags             *int                     `json:"flags,omitempty"`
	ReferencedMessage *Message                 `json:"referenced_message,omitempty"`
	Interaction       *MessageInteraction      `json:"interaction,omitempty"`
	Thread            *Channel                 `json:"thread,omitempty"`
	StickerItems      []Sticker                `json:"sticker_items,omitempty"`
	Position          *int                     `json:"position,omitempty"`
	// Note: Components are missing here, need to combine packages to allow this due to how Go imports work.
	// Will be fixed in a future breaking change.
}

type MessageReference struct {
	MessageId Snowflake `json:"message_id,omitempty"`
	// Optional when sending a reply
	ChannelId       Snowflake `json:"channel_id,omitempty"`
	GuildId         Snowflake `json:"guild_id,omitempty"`
	FailIfNotExists bool      `json:"fail_if_not_exists,omitempty"`
}

type ChannelMention struct {
	Id      Snowflake                `json:"id"`
	GuildId Snowflake                `json:"guild_id"`
	Type    channel_type.ChannelType `json:"type"`
	Name    string                   `json:"name"`
}

type MessageActivity struct {
	Type    message_activity_type.MessageActivityType `json:"type"`
	PartyId string                                    `json:"party_id"`
}

type MessageInteraction struct {
	Id     Snowflake `json:"id"`
	Type   uint8     `json:"type"`
	Name   string    `json:"name"`
	User   User      `json:"user"`
	Member *Member   `json:"member"`
}
