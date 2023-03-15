package interactions

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
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

func (ir InteractionResponse) ToAPIGatewayResponse() events.APIGatewayProxyResponse {
	data, err := json.Marshal(ir)
	if err != nil {
		fmt.Println("Error marshalling interaction response:", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}
	}
	return events.APIGatewayProxyResponse{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            string(data),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
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

type ResponseEditData struct {
	Content         *string                       `json:"content,omitempty"`
	Embeds          []discord.Embed               `json:"embeds,omitempty"`
	AllowedMentions *discord.AllowedMentions      `json:"allowed_mentions"`
	Components      []components.MessageComponent `json:"components"`
}

type AutocompleteCallbackData struct {
	Choices []commands.AutoCompleteChoice `json:"choices"`
}

type ModalCallback struct {
	CustomId   string                        `json:"custom_id"`
	Title      string                        `json:"title"`
	Components []components.MessageComponent `json:"components"`
}

func CreatePongResponse() InteractionResponse {
	return InteractionResponse{
		Type: PongInteractionCallbackType,
	}
}

func CreateDeferMessageResponse() InteractionResponse {
	return InteractionResponse{
		Type: DeferredChannelMessageWithSourceInteractionCallbackType,
	}
}

func CreateDeferEditResponse() InteractionResponse {
	return InteractionResponse{
		Type: DeferredUpdateMessageInteractionCallbackType,
	}
}
