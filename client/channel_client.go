package client

import (
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
