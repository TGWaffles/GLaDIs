package interactions

import (
	"github.com/tgwaffles/lambda-discord-interactions-go/components"
	"github.com/tgwaffles/lambda-discord-interactions-go/discord"
)

const (
	PongInteractionCallbackType                     InteractionCallbackType = iota + 1
	ChannelMessageWithSourceInteractionCallbackType                         = iota + 3
	DeferredChannelMessageWithSourceInteractionCallbackType
	DeferredUpdateMessageInteractionCallbackType
	UpdateMessageInteractionCallbackType
	ApplicationCommandAutocompleteResultInteractionCallbackType
	ModalInteractionCallbackType
)

type InteractionCallbackType uint8

type InteractionResponse struct {
	Type InteractionCallbackType `json:"type"`
	Data InteractionCallbackData `json:"data,omitempty"`
}

type InteractionCallbackData interface {
}

type MessageCallbackData struct {
	TTS             *bool                         `json:"tts,omitempty"`
	Content         *string                       `json:"content,omitempty"`
	Embeds          []discord.Embed               `json:"embeds,omitempty"`
	AllowedMentions *discord.AllowedMentions      `json:"allowed_mentions,omitempty"`
	Flags           *int                          `json:"flags,omitempty"`
	Components      []components.MessageComponent `json:"components,omitempty"`
	Attachments     []discord.Attachment          `json:"attachments,omitempty"`
}
