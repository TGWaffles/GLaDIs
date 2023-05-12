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

func (data ResponseEditData) Verify() error {
	if data.Content != nil && len(*data.Content) > 2000 {
		return fmt.Errorf("content cannot be longer than 2000 characters (you have %d)", len(*data.Content))
	}

	if len(data.Embeds) > 10 {
		return fmt.Errorf("too many embeds (max 10, you have %d)", len(data.Embeds))
	}

	for _, component := range data.Components {
		if err := component.Verify(); err != nil {
			return err
		}
	}

	for _, embed := range data.Embeds {
		if err := embed.Verify(); err != nil {
			return err
		}
	}

	return nil
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
