package client

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/tgwaffles/gladis/components"
	"github.com/tgwaffles/gladis/discord"
	"net/http"
)

type ChannelClient struct {
	ChannelId discord.Snowflake
	Bot       *BotClient
}

type SendMessageData struct {
	Content         *string                  `json:"content,omitempty"`
	TTS             *bool                    `json:"tts,omitempty"`
	Embeds          []discord.Embed          `json:"embeds,omitempty"`
	AllowedMentions *discord.AllowedMentions `json:"allowed_mentions,omitempty"`
	// Channel ID optional
	MessageReference *discord.MessageReference     `json:"message_reference,omitempty"`
	Components       []components.MessageComponent `json:"components,omitempty"`
	StickerIds       []discord.Snowflake           `json:"sticker_ids,omitempty"`
	Attachments      []discord.Attachment          `json:"attachments,omitempty"`
	// Only supports "SUPPRESS_EMBEDS" (1<<2) and "SUPPRESS_NOTIFICATIONS" (1<<12)
	Flags *int `json:"flags,omitempty"`
}

func (channelClient *ChannelClient) MakeRequest(discordRequest DiscordRequest) (response *http.Response, err error) {
	discordRequest.ValidateEndpoint()
	discordRequest.Endpoint = "/channels/" + channelClient.ChannelId.String() + discordRequest.Endpoint

	return channelClient.Bot.MakeRequest(discordRequest)
}

func (channelClient *ChannelClient) SendMessage(messageData SendMessageData) (*discord.Message, error) {
	returnedMessage := &discord.Message{}
	data, err := json.Marshal(messageData)
	if err != nil {
		return nil, fmt.Errorf("error marshaling message data to JSON: %w", err)
	}

	req := DiscordRequest{
		Method:         "POST",
		Endpoint:       "/messages",
		Body:           data,
		ExpectedStatus: 200,
		UnmarshalTo:    returnedMessage,
	}
	_, err = channelClient.MakeRequest(req)
	if err != nil {
		return nil, err
	}
	return returnedMessage, nil
}

type ThreadType uint8

type CreateThreadData struct {
	Name string `json:"name"`

	// Seconds
	AutoArchiveDuration *int `json:"auto_archive_duration,omitempty"`

	//  Valid types: AnnouncementThread, PublicThread, PrivateThread
	Type discord.ChannelType `json:"type"`

	// Whether non-mods can invite (only if private thread)
	Invitable *bool `json:"invitable,omitempty"`

	// Seconds
	RateLimit *int `json:"rate_limit_per_user,omitempty"`

	Reason *string `json:"-"` // Audit log reason, optional
}

func (channelClient *ChannelClient) CreateThread(threadData CreateThreadData) (*discord.Channel, error) {
	channel := &discord.Channel{}
	data, err := json.Marshal(threadData)
	if err != nil {
		return nil, fmt.Errorf("error marshaling thread creation data to JSON: %w", err)
	}

	additionalHeaders := make(map[string]string)

	if threadData.Reason != nil {
		additionalHeaders["X-Audit-Log-Reason"] = *threadData.Reason
	}

	req := DiscordRequest{
		Method:            "POST",
		Endpoint:          "/threads",
		Body:              data,
		ExpectedStatus:    201,
		UnmarshalTo:       channel,
		AdditionalHeaders: additionalHeaders,
	}
	_, err = channelClient.MakeRequest(req)
	if err != nil {
		return nil, err
	}
	return channel, nil
}

type ModifyChannelData interface {
	ToJson() ([]byte, error)
	GetReason() *string
}

type ModifyGroupDMChannelData struct {
	Name      *string `json:"name,omitempty"`
	IconBytes []byte  `json:"-"`
	Icon      *string `json:"icon,omitempty"`

	Reason *string `json:"-"` // Audit log reason, optional
}

func (modifyData ModifyGroupDMChannelData) ToJson() ([]byte, error) {
	if modifyData.IconBytes == nil {
		return json.Marshal(modifyData)
	}

	// Encode the icon bytes as base64
	iconAsString := base64.StdEncoding.EncodeToString(modifyData.IconBytes)
	modifyData.Icon = &iconAsString
	return json.Marshal(modifyData)
}

func (modifyData ModifyGroupDMChannelData) GetReason() *string {
	return modifyData.Reason
}

type ModifyGuildChannelData struct {
	Name *string `json:"name,omitempty"`
	// Can only switch between text and announcement channels
	Type     *discord.ChannelType `json:"type,omitempty"`
	Position *int                 `json:"position,omitempty"`
	Topic    *string              `json:"topic,omitempty"`
	NSFW     *bool                `json:"nsfw,omitempty"`
	// Slow-mode (in seconds)
	RateLimit *int `json:"rate_limit_per_user,omitempty"`

	// Voice Channels Only:
	Bitrate          *int    `json:"bitrate,omitempty"`
	UserLimit        *int    `json:"user_limit,omitempty"`
	RtcRegion        *string `json:"rtc_region,omitempty"`
	VideoQualityMode *int    `json:"video_quality_mode,omitempty"`

	// Default archive duration for user-created threads (in minutes), supports: 60, 1440, 4320, 10080
	DefaultAutoArchiveDuration *int `json:"default_auto_archive_duration,omitempty"`
	// Slow-mode (in seconds)
	DefaultThreadRateLimit *int                `json:"default_thread_rate_limit_per_user,omitempty"`
	PermissionOverwrites   []discord.Overwrite `json:"permission_overwrites,omitempty"`
	// Category ID
	ParentId *discord.Snowflake `json:"parent_id,omitempty"`

	// Forums Only:
	// Only supports REQUIRE_TAG (1<<4)
	Flags                *int                     `json:"flags,omitempty"`
	AvailableTags        []discord.Tag            `json:"available_tags,omitempty"`
	DefaultReactionEmoji *discord.DefaultReaction `json:"default_reaction_emoji,omitempty"`
	DefaultSortOrder     *int                     `json:"default_sort_order,omitempty"`
	DefaultForumLayout   *int                     `json:"default_forum_layout,omitempty"`

	Reason *string `json:"-"` // Audit log reason, optional
}

func (modifyData ModifyGuildChannelData) ToJson() ([]byte, error) {
	return json.Marshal(modifyData)
}

func (modifyData ModifyGuildChannelData) GetReason() *string {
	return modifyData.Reason
}

type ModifyThreadData struct {
	Name     *string `json:"name,omitempty"`
	Archived *bool   `json:"archived,omitempty"`
	// Minutes, only supports: 60, 1440, 4320, 10080
	AutoArchiveDuration *int  `json:"auto_archive_duration,omitempty"`
	Locked              *bool `json:"locked,omitempty"`
	Invitable           *bool `json:"invitable,omitempty"`
	// Slow-mode (in seconds)
	RateLimit *int `json:"rate_limit_per_user,omitempty"`

	Flags       *int          `json:"flags,omitempty"`
	AppliedTags []discord.Tag `json:"applied_tags,omitempty"`

	Reason *string `json:"-"` // Audit log reason, optional
}

func (modifyData ModifyThreadData) ToJson() ([]byte, error) {
	return json.Marshal(modifyData)
}

func (modifyData ModifyThreadData) GetReason() *string {
	return modifyData.Reason
}

func (channelClient *ChannelClient) Edit(data ModifyChannelData) (*discord.Channel, error) {
	channel := &discord.Channel{}
	jsonData, err := data.ToJson()
	if err != nil {
		return nil, fmt.Errorf("error marshaling channel modification data to JSON: %w", err)
	}

	additionalHeaders := make(map[string]string)

	if data.GetReason() != nil {
		additionalHeaders["X-Audit-Log-Reason"] = *data.GetReason()
	}

	req := DiscordRequest{
		Method:         "PATCH",
		Endpoint:       "",
		Body:           jsonData,
		ExpectedStatus: 200,
		UnmarshalTo:    channel,
	}
	_, err = channelClient.MakeRequest(req)
	if err != nil {
		return nil, err
	}
	return channel, nil
}
