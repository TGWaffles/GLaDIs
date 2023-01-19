package interactions

import (
	"github.com/tgwaffles/gladis/commands"
	"github.com/tgwaffles/gladis/components"
	"github.com/tgwaffles/gladis/discord"
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

type AutocompleteCallbackData struct {
	Choices []commands.AutoCompleteChoice `json:"choices"`
}

type ModalCallback struct {
	CustomId   string                        `json:"custom_id"`
	Title      string                        `json:"title"`
	Components []components.MessageComponent `json:"components"`
}
